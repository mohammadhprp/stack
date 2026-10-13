import copy
import json
import os
import re
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path
from unittest.mock import patch


ROOT = Path(__file__).resolve().parents[2]
SCRIPTS = ROOT / "framework/skills/walkthrough/scripts"
sys.path.insert(0, str(SCRIPTS))

import render_walkthrough as renderer
from validate_walkthrough_json import validate


def example():
    schema = (SCRIPTS.parent / "references/schema.md").read_text()
    return json.loads(schema.split("```json\n", 1)[1].split("```", 1)[0])


class WalkthroughTests(unittest.TestCase):
    def test_example_and_all_labels(self):
        data = example()
        for kind in ("Warning", "Notice", "Human decision"):
            item = copy.deepcopy(data["groups"][0]["attention"][0])
            item["kind"] = kind
            data["groups"][0]["attention"] = [item]
            validate(data)
            self.assertIn(kind, renderer.markdown(data))
            self.assertIn(kind, renderer.render_html(data))

    def test_optional_mermaid_and_highlighted_code(self):
        data = example()
        group = data["groups"][0]
        group["diagrams"] = [
            {
                "title": "Payment flow",
                "source": 'flowchart TD\n  A["<script>"] --> B{Retry?}\n  B -->|yes| C[Reuse]',
            }
        ]
        group["code"] = [
            {
                "title": "Retry guard",
                "language": "go",
                "source": 'if retry {\n    return "existing" // keep prior payment\n}',
            }
        ]
        validate(data)
        html = renderer.render_html(data)
        markdown = renderer.markdown(data)
        self.assertIn("mermaid@12.2.0", html)
        self.assertIn("securityLevel:'strict'", html)
        self.assertIn(
            "Diagram source; rendering requires an internet connection.", html
        )
        self.assertIn("&lt;script&gt;", html)
        self.assertIn('class="tok-keyword">if</span>', html)
        self.assertIn('class="tok-string">&quot;existing&quot;</span>', html)
        self.assertIn('class="tok-comment">// keep prior payment</span>', html)
        self.assertIn("```mermaid\nflowchart TD", markdown)
        self.assertIn("```go\nif retry", markdown)

        group.pop("diagrams")
        html_without_diagram = renderer.render_html(data)
        self.assertNotIn("<script", html_without_diagram)
        self.assertNotIn("cdn.jsdelivr.net", html_without_diagram)

    def test_malformed_optional_fields_are_rejected(self):
        for key, value in (
            ("diagrams", [{}]),
            ("code", [{"title": "Example", "language": "go", "source": "\U0001f6a8"}]),
        ):
            with self.subTest(key=key):
                data = example()
                data["groups"][0][key] = value
                with self.assertRaises(ValueError):
                    validate(data)

    def test_validation_rejects_invalid_content(self):
        for change in (
            lambda d: d.update(title=""),
            lambda d: d.update(title="Warning \U0001f6a8"),
            lambda d: d["groups"][0].update(category="Unknown"),
            lambda d: d["groups"][0]["attention"][0].update(kind="Critical"),
            lambda d: d["groups"][0]["attention"][0].update(evidence=""),
            lambda d: d["groups"][0].update(pieces=[]),
            lambda d: d["changed_files"].append("missing.go"),
            lambda d: d["changed_files"].append("../escape.go"),
            lambda d: d["changed_files"].append("/absolute.go"),
            lambda d: d["changed_files"].append(d["changed_files"][0]),
            lambda d: d["checklist"][0].update(group=2),
            lambda d: d["checklist"][0].update(group=True),
            lambda d: d.update(limitations="unknown"),
        ):
            with self.subTest(change=change):
                data = example()
                change(data)
                with self.assertRaises(ValueError):
                    validate(data)

    def test_empty_diff(self):
        data = example()
        data.update(changed_files=[], groups=[], checklist=[], limitations=[])
        validate(data)
        for report in (renderer.markdown(data), renderer.render_html(data)):
            self.assertIn("No changes in the reviewed scope.", report)
            self.assertIn("None.", report)

    def test_group_order_links_and_escaping(self):
        data = example()
        data["groups"].append(copy.deepcopy(data["groups"][0]))
        data["groups"][1]["title"] = "Second group"
        data["groups"][1]["attention"] = []
        data["title"] = '<script>alert("unsafe")</script>'
        validate(data)
        html = renderer.render_html(data)
        markdown = renderer.markdown(data)
        self.assertNotIn("<script>", html)
        self.assertNotIn("<script>", markdown)
        self.assertIn("&lt;script&gt;", html)
        self.assertIn("None identified.", html)
        self.assertIn("color-scheme:dark", html)
        self.assertIn("max-width:650px", html)
        self.assertIn("@media print", html)
        self.assertNotIn("<details", html)
        self.assertNotIn("<link", html)
        targets = set(re.findall(r'id="([^"]+)"', html))
        for target in re.findall(r'href="#([^"]+)"', html):
            self.assertIn(target, targets)
        self.assertLess(html.index('id="group-1"'), html.index('id="group-2"'))
        self.assertIn('<a id="group-2"></a>', markdown)
        for group in data["groups"]:
            for value in (group["summary"], group["rationale"], group["verification"]):
                self.assertIn(value, html)
                self.assertIn(value, markdown)

    def run_renderer(self, *args, **kwargs):
        return subprocess.run(
            [sys.executable, "-B", str(SCRIPTS / "render_walkthrough.py"), *args],
            capture_output=True,
            text=True,
            **kwargs,
        )

    def test_cli_defaults_and_safe_rerun(self):
        with tempfile.TemporaryDirectory(prefix="walkthrough ") as directory:
            root = Path(directory)
            project = root / "demo project"
            project.mkdir()
            home = root / "home"
            output = home / ".stack/demo project/walkthrough"
            output.mkdir(parents=True)
            (output / "walkthrough.json").write_text(json.dumps(example()))
            env = dict(os.environ, HOME=str(home))
            result = self.run_renderer(cwd=project, env=env)
            self.assertEqual(result.returncode, 0, result.stderr)
            for name in ("walkthrough.md", "walkthrough.html"):
                self.assertTrue((output / name).exists())
                self.assertIn(str(output / name), result.stdout)
                self.assertFalse((project / name).exists())
            before = (output / "walkthrough.md").read_text()
            again = self.run_renderer(cwd=project, env=env)
            self.assertNotEqual(again.returncode, 0)
            self.assertIn("already exists", again.stderr)
            self.assertEqual((output / "walkthrough.md").read_text(), before)
            forced = self.run_renderer("--force", cwd=project, env=env)
            self.assertEqual(forced.returncode, 0, forced.stderr)

    def test_existing_html_does_not_create_markdown(self):
        with tempfile.TemporaryDirectory() as directory:
            output = Path(directory)
            source = output / "walkthrough.json"
            source.write_text(json.dumps(example()))
            (output / "walkthrough.html").write_text("user content")
            result = self.run_renderer("--output-dir", str(output))
            self.assertNotEqual(result.returncode, 0)
            self.assertFalse((output / "walkthrough.md").exists())
            self.assertEqual((output / "walkthrough.html").read_text(), "user content")

    def test_invalid_input_does_not_create_output_dir(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            source = root / "bad.json"
            source.write_text("{}")
            output = root / "reports"
            result = self.run_renderer(
                "--walkthrough-json", str(source), "--output-dir", str(output)
            )
            self.assertNotEqual(result.returncode, 0)
            self.assertFalse(output.exists())

    def test_git_root_and_no_git_fallback(self):
        with tempfile.TemporaryDirectory() as directory:
            project = Path(directory) / "project"
            subdir = project / "src"
            subdir.mkdir(parents=True)
            subprocess.run(["git", "init", "-q", str(project)], check=True)
            self.assertEqual(renderer.default_output_dir(subdir).name, "walkthrough")
            self.assertEqual(renderer.default_output_dir(subdir).parent.name, "project")
            with patch.object(
                renderer.subprocess, "run", side_effect=FileNotFoundError
            ):
                self.assertEqual(renderer.default_output_dir(subdir).parent.name, "src")


if __name__ == "__main__":
    unittest.main()
