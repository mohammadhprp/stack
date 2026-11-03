import assert from "node:assert/strict";
import { mkdtempSync, mkdirSync, rmSync, symlinkSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { spawnSync } from "node:child_process";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { MARKER_END, MARKER_START, replaceMarkedBlock } from "../../framework/skills/before-and-after/scripts/format.mjs";

const script = fileURLToPath(new URL("../../framework/skills/before-and-after/scripts/format.mjs", import.meta.url));

test("media references stay inside the real working directory", (t) => {
  const dir = mkdtempSync(join(tmpdir(), "stack-media-"));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  const cwd = join(dir, "output");
  mkdirSync(cwd);
  writeFileSync(join(cwd, "preview.png"), "image");
  writeFileSync(join(dir, "private.png"), "private");
  symlinkSync(join(dir, "private.png"), join(cwd, "escape.png"));
  symlinkSync(cwd, join(dir, "output-link"));
  mkdirSync(join(cwd, "directory.png"));

  for (const workingDir of [cwd, join(dir, "output-link")]) {
    const result = spawnSync(process.execPath, [script, "--after", "preview.png", "--attach-list"], { cwd: workingDir, encoding: "utf8" });
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, "./preview.png\n");
  }
  for (const file of ["../private.png", "escape.png", "directory.png"]) {
    const result = spawnSync(process.execPath, [script, "--after", file, "--attach-list"], { cwd, encoding: "utf8" });
    assert.equal(result.status, 1, `${file}: ${result.stderr}`);
    assert.equal(result.stdout, "");
  }
});

test("replacing media preserves all bytes outside its markers", () => {
  const before = "Title\r\n\r\n ";
  const after = " \r\nReviewer notes\r\n";
  const block = `${MARKER_START}\nnew\n${MARKER_END}\n`;
  const body = `${before}${MARKER_START}\nold\n${MARKER_END}${after}`;
  const expected = before + block.slice(0, -1) + after;
  assert.equal(replaceMarkedBlock(body, block), expected);
  assert.equal(replaceMarkedBlock(expected, block), expected);
});
