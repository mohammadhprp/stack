#!/usr/bin/env python3

import argparse
import html
import re
import subprocess
import sys
from pathlib import Path

sys.dont_write_bytecode = True

from validate_walkthrough_json import load


CSS = """
:root { color-scheme:dark; --bg:#121212; --panel:#1b1b1b; --ink:#f4f4f0;
  --muted:#a9aaa5; --line:#3b3b3b; --blue:#8cc8ff; --red:#ff7d7d;
  --orange:#ffbf69; --green:#78dba9; --purple:#c7a6ff; --warning:#ffbf69;
  --notice:#8cc8ff; --decision:#c7a6ff; }
* { box-sizing:border-box; }
body { margin:0; background:var(--bg); color:var(--ink); font:16px/1.55 ui-rounded,"Comic Sans MS",system-ui,sans-serif; }
body:before { content:""; position:fixed; inset:0; pointer-events:none; opacity:.08; background-image:radial-gradient(#fff .7px,transparent .7px); background-size:13px 13px; }
main { width:min(1060px,calc(100% - 32px)); margin:48px auto; }
.eyebrow,.category { color:var(--blue); font:700 12px/1 monospace; letter-spacing:.16em; text-transform:uppercase; }
h1 { font-size:clamp(32px,6vw,62px); line-height:1; margin:10px 0 28px; letter-spacing:-.06em; }
h2 { font-size:22px; margin:0 0 14px; }
h3 { font-size:16px; margin:22px 0 8px; }
p { margin:8px 0 14px; white-space:pre-wrap; overflow-wrap:anywhere; color:#e4e4df; }
a { color:var(--blue); text-underline-offset:3px; }
a:focus-visible { outline:3px solid var(--blue); outline-offset:4px; }
.hero,.panel { background:var(--panel); border:2px solid var(--line); border-radius:13px 10px 15px 9px; box-shadow:5px 5px 0 #080808; }
.hero { padding:24px; margin-top:22px; }
.metadata { display:grid; grid-template-columns:100px 1fr; gap:6px 12px; font-size:14px; }
.metadata dt { color:var(--muted); } .metadata dd { margin:0; overflow-wrap:anywhere; }
.stats { display:grid; grid-template-columns:repeat(4,1fr); gap:12px; margin:22px 0; }
.stat { padding:16px; border:2px solid var(--line); border-radius:8px 11px 7px 10px; background:#171717; }
.stat b { display:block; font-size:30px; line-height:1; } .stat span { color:var(--muted); font-size:12px; text-transform:uppercase; letter-spacing:.08em; }
.panel { padding:22px; margin:22px 0; }
nav ol,ol.reading { padding-left:24px; } li { margin:10px 0; overflow-wrap:anywhere; }
.group { scroll-margin-top:24px; } .category { display:block; margin-bottom:10px; }
.comparison { display:grid; grid-template-columns:1fr 1fr; gap:12px; }
.comparison > div { border:2px solid var(--line); border-radius:8px 11px 7px 10px; padding:0 16px; background:#171717; }
code { font:13px/1.6 ui-monospace,SFMono-Regular,Consolas,monospace; overflow-wrap:anywhere; }
.location,.muted { color:var(--muted); } .role { display:block; color:#e4e4df; }
.attention { border-left:5px solid var(--blue); border-radius:4px 10px 8px 5px; padding:14px 16px 16px; margin:12px 0; background:#151515; }
.attention.warning { border-color:var(--orange); } .attention.notice { border-color:var(--blue); } .attention.decision { border-color:var(--purple); }
.attention p:last-child { margin-bottom:0; }
.attention-label { color:var(--muted); font:700 11px monospace; letter-spacing:.08em; text-transform:uppercase; margin-bottom:7px; }
.attention.warning .attention-label { color:var(--orange); } .attention.notice .attention-label { color:var(--blue); }
.attention.decision .attention-label { color:var(--purple); }
.code-sample,.diagram { margin:18px 0; }
.code-sample figcaption,.diagram figcaption { color:var(--muted); font:700 12px monospace; letter-spacing:.06em; margin-bottom:8px; }
.code-sample figcaption span { color:var(--blue); text-transform:uppercase; }
.codeblock,.diagram-source { overflow:auto; margin:10px 0 0; padding:13px; background:#0c0c0c; border:1px solid var(--line); border-radius:7px; color:#d7d7d0; font:13px/1.6 ui-monospace,SFMono-Regular,Consolas,monospace; white-space:pre; }
.tok-keyword { color:var(--purple); } .tok-string { color:#a8d8a8; } .tok-number { color:var(--orange); } .tok-comment { color:#777b78; font-style:italic; }
.diagram-canvas { overflow:auto; padding:12px; background:#151515; border:1px solid var(--line); border-radius:7px; }
.diagram-canvas svg { max-width:100%; height:auto; } .diagram-status { color:var(--muted); font-size:12px; }
@media (max-width:650px) { main { margin:24px auto; } .stats { grid-template-columns:repeat(2,1fr); }
  .comparison { grid-template-columns:1fr; } .hero { padding:18px; } }
@media print { body { background:white; color:#111; } body:before { display:none; }
  main { width:100%; margin:0; } .hero,.panel { box-shadow:none; break-inside:avoid; }
  .diagram-canvas { display:none; } .diagram-source { display:block !important; } }
"""


def esc(value):
    return html.escape(value, quote=True)


def md(value):
    value = html.escape(value, quote=False)
    return re.sub(r"([\\`*_\[\]#])", r"\\\1", value)


CODE_TOKEN = re.compile(
    r"(?P<comment>//[^\n]*|\#[^\n]*|/\*[\s\S]*?\*/)|"
    r"(?P<string>\"(?:\\.|[^\"\\])*\"|'(?:\\.|[^'\\])*'|`(?:\\.|[^`\\])*`)|"
    r"(?P<keyword>\b(?:const|let|var|function|return|if|else|for|while|class|new|throw|async|await|import|from|export|type|interface|func|package|var|go|defer|range|select|switch|case|default|break|continue|try|catch|finally|def|lambda|pass|yield|with|as|in|is|not|and|or|true|false|null|None|nil|undefined)\b)|"
    r"(?P<number>\b\d+(?:\.\d+)?\b)"
)


def highlighted_code(source):
    parts = []
    last = 0
    classes = {
        "comment": "tok-comment",
        "string": "tok-string",
        "keyword": "tok-keyword",
        "number": "tok-number",
    }
    for match in CODE_TOKEN.finditer(source):
        parts.append(esc(source[last : match.start()]))
        parts.append(
            f'<span class="{classes[match.lastgroup]}">{esc(match.group())}</span>'
        )
        last = match.end()
    parts.append(esc(source[last:]))
    return "".join(parts)


def fence(source):
    longest = max((len(run) for run in re.findall(r"`+", source)), default=2)
    return "`" * max(3, longest + 1)


def markdown(data):
    lines = [
        f"# Walkthrough: {md(data['title'])}",
        "",
        "## Overview",
        "",
        f"Source: {md(data['source'])}",
        f"Revisions: {md(data['revisions'])}",
        "",
        md(data["overview"]),
        "",
        "## Change map",
        "",
    ]
    for index, group in enumerate(data["groups"], 1):
        lines.append(
            f"{index}. [{md(group['category'])} - {md(group['title'])}](#group-{index})"
        )
    if not data["groups"]:
        lines.append("No changes in the reviewed scope.")
    for index, group in enumerate(data["groups"], 1):
        lines += [
            "",
            f'<a id="group-{index}"></a>',
            "",
            f"## {index}. {md(group['category'])} - {md(group['title'])}",
            "",
            "### What changed",
            "",
            md(
                group["summary"]
                + (" " + group["rationale"] if group["rationale"] else "")
            ),
            "",
            "### Before and after",
            "",
            f"**Before:** {md(group['before'])}",
            "",
            f"**After:** {md(group['after'])}",
            "",
            "### Read these pieces together",
            "",
        ]
        for number, piece in enumerate(group["pieces"], 1):
            lines.append(
                f"{number}. {md(piece['path'])} ({md(piece['location'])}) - {md(piece['role'])}"
            )
        lines += [
            "",
            "### Verification",
            "",
            md(group["verification"]),
            "",
        ]
        if group.get("diagrams") or group.get("code"):
            lines += ["### Visual aids", ""]
        for diagram in group.get("diagrams", []):
            lines += [
                f"### {md(diagram['title'])}",
                "",
                f"```mermaid\n{diagram['source']}\n```",
                "",
            ]
        for snippet in group.get("code", []):
            marker = fence(snippet["source"])
            lines += [
                f"### {md(snippet['title'])}",
                "",
                f"{marker}{snippet['language']}",
                snippet["source"],
                marker,
                "",
            ]
        lines += ["### Human attention", ""]
        if not group["attention"]:
            lines.append("None identified.")
        for item in group["attention"]:
            lines += [
                f"**{item['kind']}:** {md(item['message'])}",
                "",
                f"Evidence: {md(item['evidence'])}",
                "",
                f"Action: {md(item['action'])}",
                "",
            ]
    lines += ["", "## Reviewer checklist", ""]
    lines += [
        f"- [ ] {md(item['text'])} ([Group {item['group']}](#group-{item['group']}))"
        for item in data["checklist"]
    ] or ["None identified."]
    lines += ["", "## Limitations", ""]
    lines += [f"- {md(item)}" for item in data["limitations"]] or ["None."]
    return "\n".join(lines).rstrip() + "\n"


def render_html(data):
    counts = {kind: 0 for kind in ("Warning", "Notice", "Human decision")}
    for group in data["groups"]:
        for item in group["attention"]:
            counts[item["kind"]] += 1
    has_diagrams = any(group.get("diagrams") for group in data["groups"])
    parts = [
        '<!doctype html><html lang="en"><head><meta charset="utf-8">',
        '<meta name="viewport" content="width=device-width,initial-scale=1">',
        f"<title>Walkthrough: {esc(data['title'])}</title><style>{CSS}</style></head><body><main>",
        '<div class="eyebrow">walkthrough artifact // visualized</div>',
        f"<h1>{esc(data['title'])}</h1>",
        '<section class="hero"><div class="eyebrow">change overview</div><dl class="metadata">',
        f"<dt>Source</dt><dd>{esc(data['source'])}</dd>",
        f"<dt>Revisions</dt><dd>{esc(data['revisions'])}</dd></dl>",
        f"<p>{esc(data['overview'])}</p></section>",
        '<section class="stats">',
        f'<div class="stat"><b>{len(data["groups"])}</b><span>Change groups</span></div>',
        f'<div class="stat"><b>{counts["Warning"]}</b><span>Warnings</span></div>',
        f'<div class="stat"><b>{counts["Notice"]}</b><span>Notices</span></div>',
        f'<div class="stat"><b>{counts["Human decision"]}</b><span>Human decisions</span></div>',
        "</section>",
        '<section class="panel"><h2>Change map</h2><ol>',
    ]
    for index, group in enumerate(data["groups"], 1):
        parts.append(
            f'<li><a href="#group-{index}">{esc(group["category"])} - {esc(group["title"])}</a></li>'
        )
    parts.append(
        "</ol>"
        + ("" if data["groups"] else "<p>No changes in the reviewed scope.</p>")
        + "</section>"
    )
    styles = {"Warning": "warning", "Notice": "notice", "Human decision": "decision"}
    for index, group in enumerate(data["groups"], 1):
        parts += [
            f'<section class="panel group" id="group-{index}" aria-labelledby="title-{index}">',
            f'<span class="category">{esc(group["category"])}</span>',
            f'<h2 id="title-{index}">{index}. {esc(group["title"])}</h2>',
            f"<h3>What changed</h3><p>{esc(group['summary'])} {esc(group['rationale'])}</p>",
            '<h3>Before and after</h3><div class="comparison">',
            f"<div><h3>Before</h3><p>{esc(group['before'])}</p></div>",
            f"<div><h3>After</h3><p>{esc(group['after'])}</p></div></div>",
            '<h3>Read these pieces together</h3><ol class="reading">',
        ]
        for piece in group["pieces"]:
            parts.append(
                f"<li><code>{esc(piece['path'])}</code> "
                f'<span class="location">({esc(piece["location"])})</span>'
                f'<span class="role">{esc(piece["role"])}</span></li>'
            )
        parts += [
            f"</ol><h3>Verification</h3><p>{esc(group['verification'])}</p>",
        ]
        if group.get("diagrams") or group.get("code"):
            parts.append("<h3>Visual aids</h3>")
        for diagram_index, diagram in enumerate(group.get("diagrams", []), 1):
            parts += [
                f'<figure class="diagram" data-diagram="{index}-{diagram_index}">',
                f"<figcaption>{esc(diagram['title'])}</figcaption>",
                f'<div class="diagram-canvas" aria-label="{esc(diagram["title"])}"></div>',
                f'<pre class="diagram-source"><code>{esc(diagram["source"])}</code></pre>',
                '<p class="diagram-status" role="status">Diagram source; rendering requires an internet connection.</p>',
                "</figure>",
            ]
        for snippet in group.get("code", []):
            parts += [
                '<figure class="code-sample">',
                f"<figcaption>{esc(snippet['title'])} <span>{esc(snippet['language'])}</span></figcaption>",
                f'<pre class="codeblock"><code>{highlighted_code(snippet["source"])}</code></pre>',
                "</figure>",
            ]
        parts.append("<h3>Human attention</h3>")
        if not group["attention"]:
            parts.append('<p class="muted">None identified.</p>')
        for item in group["attention"]:
            parts.append(
                f'<div class="attention {styles[item["kind"]]}">'
                f'<div class="attention-label">{item["kind"]}</div>'
                f"<p>{esc(item['message'])}</p><p><strong>Evidence:</strong> {esc(item['evidence'])}</p>"
                f"<p><strong>Action:</strong> {esc(item['action'])}</p></div>"
            )
        parts.append("</section>")
    parts.append('<section class="panel"><h2>Reviewer checklist</h2><ul>')
    for item in data["checklist"]:
        parts.append(
            f'<li>{esc(item["text"])} <a href="#group-{item["group"]}">Group {item["group"]}</a></li>'
        )
    parts.append(
        "</ul>"
        + ("" if data["checklist"] else "<p>None identified.</p>")
        + "</section>"
    )
    parts.append('<section class="panel"><h2>Limitations</h2>')
    parts.append(
        "<ul>"
        + "".join(f"<li>{esc(item)}</li>" for item in data["limitations"])
        + "</ul>"
        if data["limitations"]
        else "<p>None.</p>"
    )
    if has_diagrams:
        parts.append(
            """<script type="module">
import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@12.2.0/dist/mermaid.esm.min.mjs';
mermaid.initialize({startOnLoad:false,securityLevel:'strict',theme:'dark'});
for (const figure of document.querySelectorAll('.diagram')) {
  const source = figure.querySelector('.diagram-source');
  const status = figure.querySelector('.diagram-status');
  try {
    const {svg} = await mermaid.render(`walkthrough-${figure.dataset.diagram}`, source.textContent);
    figure.querySelector('.diagram-canvas').innerHTML = svg;
    source.hidden = true;
    status.hidden = true;
  } catch (_) {
    status.textContent = 'Diagram could not be rendered. Mermaid source is shown below.';
  }
}
</script>"""
        )
    parts.append("</section></main></body></html>\n")
    return "\n".join(parts)


def default_output_dir(project):
    root = Path(project).expanduser().resolve()
    try:
        result = subprocess.run(
            ["git", "-C", str(root), "rev-parse", "--show-toplevel"],
            capture_output=True,
            text=True,
            check=False,
        )
    except FileNotFoundError:
        result = None
    if result is not None and result.returncode == 0:
        root = Path(result.stdout.strip())
    if not root.name:
        raise ValueError("Cannot determine a project name; pass --output-dir")
    return Path.home() / ".stack" / root.name / "walkthrough"


def main():
    parser = argparse.ArgumentParser(
        description="Render matching Markdown and styled HTML walkthroughs."
    )
    parser.add_argument("--walkthrough-json")
    parser.add_argument("--project-root", default=".")
    parser.add_argument("--output-dir")
    parser.add_argument(
        "--force", action="store_true", help="Overwrite existing rendered outputs"
    )
    args = parser.parse_args()
    try:
        output = (
            Path(args.output_dir).expanduser()
            if args.output_dir
            else default_output_dir(args.project_root)
        ).resolve()
        source = (
            Path(args.walkthrough_json).expanduser()
            if args.walkthrough_json
            else output / "walkthrough.json"
        )
        data = load(source)
        files = {
            output / "walkthrough.md": markdown(data),
            output / "walkthrough.html": render_html(data),
        }
        for file in files:
            if file.exists() and not args.force:
                raise ValueError(
                    f"Output already exists: {file}; use --force only with overwrite approval"
                )
            if file.resolve() == source.resolve():
                raise ValueError("Input path must not be a rendered output path")
        output.mkdir(parents=True, exist_ok=True)
        for file, content in files.items():
            with file.open("w" if args.force else "x", encoding="utf-8") as stream:
                stream.write(content)
            print(file)
    except (ValueError, OSError) as error:
        parser.exit(1, f"Cannot render walkthrough: {error}\n")


if __name__ == "__main__":
    main()
