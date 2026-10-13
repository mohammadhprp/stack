#!/usr/bin/env python3

import argparse
import json
import re
from pathlib import Path, PurePosixPath


CATEGORIES = {"Feature", "Bug fix", "Refactor", "Tests", "Documentation", "Maintenance"}
ATTENTION_KINDS = {"Warning", "Notice", "Human decision"}
EMOJI = re.compile("[\U0001f000-\U0001faff\u2600-\u27bf\ufe0f\u20e3]")


def text(value, label):
    if not isinstance(value, str) or not value.strip():
        raise ValueError(f"{label} must be a non-empty string")
    if EMOJI.search(value):
        raise ValueError(f"{label} must not contain emoji")


def array(value, label):
    if not isinstance(value, list):
        raise ValueError(f"{label} must be an array")


def fields(value, names, label):
    if not isinstance(value, dict):
        raise ValueError(f"{label} must be an object")
    for name in names:
        text(value.get(name), f"{label}.{name}")


def path(value, label):
    text(value, label)
    parsed = PurePosixPath(value)
    if (
        parsed.is_absolute()
        or ".." in parsed.parts
        or "\\" in value
        or ":" in value
        or value == "."
    ):
        raise ValueError(f"{label} must be a repository-relative path")


def validate(data):
    fields(data, ("title", "source", "revisions", "overview"), "walkthrough")
    for key in ("changed_files", "groups", "checklist", "limitations"):
        array(data.get(key), key)
    for item in data["changed_files"]:
        path(item, "changed_files item")
    if len(set(data["changed_files"])) != len(data["changed_files"]):
        raise ValueError("changed_files must not contain duplicates")
    if data["changed_files"] and not data["groups"]:
        raise ValueError("changed files need at least one group")
    covered = set()
    for index, group in enumerate(data["groups"], 1):
        label = f"group {index}"
        fields(
            group,
            (
                "title",
                "category",
                "summary",
                "rationale",
                "before",
                "after",
                "verification",
            ),
            label,
        )
        if group["category"] not in CATEGORIES:
            raise ValueError(f"{label}: unsupported category {group['category']!r}")
        array(group.get("pieces"), f"{label}.pieces")
        array(group.get("attention"), f"{label}.attention")
        for key in ("diagrams", "code"):
            if key in group:
                array(group[key], f"{label}.{key}")
        if not group["pieces"]:
            raise ValueError(f"{label} needs at least one reading piece")
        for piece in group["pieces"]:
            fields(piece, ("path", "location", "role"), f"{label}.piece")
            path(piece["path"], f"{label}.piece.path")
            covered.add(piece["path"])
        for item in group["attention"]:
            fields(
                item, ("kind", "message", "evidence", "action"), f"{label}.attention"
            )
            if item["kind"] not in ATTENTION_KINDS:
                raise ValueError(
                    f"{label}: unsupported attention kind {item['kind']!r}"
                )
        for diagram in group.get("diagrams", []):
            fields(diagram, ("title", "source"), f"{label}.diagram")
        for snippet in group.get("code", []):
            fields(snippet, ("title", "language", "source"), f"{label}.code")
    missing = set(data["changed_files"]) - covered
    if missing:
        raise ValueError(
            f"changed files missing from reading pieces: {', '.join(sorted(missing))}"
        )
    for item in data["checklist"]:
        fields(item, ("text",), "checklist item")
        group = item.get("group")
        if type(group) is not int or not 1 <= group <= len(data["groups"]):
            raise ValueError("checklist group must reference a group number")
    for item in data["limitations"]:
        text(item, "limitation")
    return data


def load(file):
    return validate(json.loads(Path(file).expanduser().read_text(encoding="utf-8")))


def main():
    parser = argparse.ArgumentParser(
        description="Validate walkthrough content and changed-file coverage."
    )
    parser.add_argument("--walkthrough-json", required=True)
    args = parser.parse_args()
    try:
        load(args.walkthrough_json)
    except (ValueError, OSError) as error:
        parser.exit(1, f"Invalid walkthrough: {error}\n")
    print("Valid walkthrough JSON; all changed files are grouped.")


if __name__ == "__main__":
    main()
