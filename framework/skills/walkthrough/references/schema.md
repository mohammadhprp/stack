# Walkthrough content standard

Write `$OUTPUT_DIR/walkthrough.json` using the following shape. All fields shown
are required. Text is plain text, not HTML or Markdown; the renderer escapes it.
Do not include emoji, even in titles copied from the PR/MR: paraphrase them.

```json
{
  "title": "Prevent duplicate payments",
  "source": "Local staged and unstaged changes",
  "revisions": "Working tree against HEAD abc1234",
  "overview": "Repeated requests reuse an existing payment instead of charging again.",
  "changed_files": ["internal/payment.go", "tests/payment_test.go"],
  "groups": [
    {
      "title": "Reject duplicate payment requests",
      "category": "Bug fix",
      "summary": "The handler checks for an existing payment before creating one.",
      "rationale": "The regression test describes repeated requests as retries.",
      "before": "Each request could create a payment.",
      "after": "Sequential retries return the original payment.",
      "pieces": [
        {"path": "internal/payment.go", "location": "CreatePayment", "role": "Implements retry handling."},
        {"path": "tests/payment_test.go", "location": "TestPaymentRetry", "role": "Covers a sequential retry."}
      ],
      "verification": "Regression test inspected, not executed. No CI results available.",
      "diagrams": [
        {
          "title": "Request and retry flow",
          "source": "flowchart TD\n  Request[Payment request] --> Lookup{Existing payment?}\n  Lookup -- yes --> Reuse[Return existing payment]\n  Lookup -- no --> Create[Create payment]"
        }
      ],
      "code": [
        {
          "title": "Retry guard",
          "language": "go",
          "source": "if existing, ok := payments[key]; ok {\n    return existing, nil\n}"
        }
      ],
      "attention": [
        {
          "kind": "Warning",
          "message": "Concurrent retries remain unverified; this is not a confirmed race.",
          "evidence": "TestPaymentRetry submits requests sequentially.",
          "action": "Check the persistence constraint and concurrent-request coverage."
        }
      ]
    }
  ],
  "checklist": [{"group": 1, "text": "Verify concurrent retry behavior."}],
  "limitations": ["No runtime checks were executed."]
}
```

- Categories: `Feature`, `Bug fix`, `Refactor`, `Tests`, `Documentation`, `Maintenance`.
- Attention kinds: `Warning`, `Notice`, `Human decision`. An empty array means
  no attention items were identified; it is not an approval.
- `pieces` is the ordered reading guide. Paths are repository-relative, with
  symbols or verified line ranges in `location`. Every `changed_files` entry
  must appear in a group's pieces. Unchanged supporting files may also appear.
- Include impact in the attention message, a grounded reference in `evidence`,
  and a specific human check or question in `action`. Never invent an owner.
- Keep content skimmable: overview 1-2 sentences; group summary and rationale
  together no longer than 2 short sentences; before/after and each file role one
  short sentence; verification one concise result. Omit redundant rationale,
  routine caveats, and generic checklist items. Include only material limitations
  and checks that require a human decision.
- `diagrams` and `code` are optional arrays on each group. Add a Mermaid diagram
  only when it explains a real flow, dependency, or state change more clearly
  than prose. Keep it small and grounded in the reviewed change. Add short,
  exact code excerpts only when they clarify the walkthrough; identify the
  language for highlighting and cite their source file in `pieces`.
- Checklist `group` values are one-based group numbers used for section links.
- Use empty arrays for absent checklist items and limitations. An empty diff may
  use empty `changed_files` and `groups`; explain its scope in the overview.
- Validation checks structure, labels, emoji, links between checklist and groups,
  and changed-file coverage. It cannot verify claims or line numbers: inspect
  those against the actual diff before rendering.

The renderer owns presentation: a dark, textured page, oversized heading, bordered
overview card, count tiles, linked change map, section cards, side-by-side before/after
blocks, ordered reading pieces, optional Mermaid diagrams and highlighted code,
labeled attention cards, and concise checklist and limitations when needed. The visual language matches the `review` report: near-black surfaces,
compact monospace labels, irregular rounded borders, subtle hard shadows, and blue,
orange, and purple accents. Its self-contained HTML uses visible keyboard focus,
responsive layout, and print styling. All sections remain visible; there is no wizard,
approval badge, or emoji. If a Mermaid diagram is present, only the pinned Mermaid
ESM renderer is loaded from jsDelivr; the report displays its source if the CDN
is unavailable. Without diagrams, no remote assets or scripts are included.
