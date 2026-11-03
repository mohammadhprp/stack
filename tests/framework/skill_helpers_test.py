import importlib.util
import io
import json
import re
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest
import zipfile
from unittest.mock import patch
from urllib.error import HTTPError, URLError


SKILLS = Path(__file__).resolve().parents[2] / "framework" / "skills"


class SkillHelpersTest(unittest.TestCase):
    def test_eval_viewer_preserves_text_without_script_injection(self):
        spec = importlib.util.spec_from_file_location("eval_viewer", SKILLS / "skill-creator/eval-viewer/generate_review.py")
        viewer = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(viewer)
        payload = "</script><script>alert('private')</script>"
        runs = [{"id": "test", "prompt": payload, "outputs": []}]
        html = viewer.generate_html(runs, "sample")
        self.assertNotIn(payload, html)
        embedded = re.search(r"const EMBEDDED_DATA = (.*);", html).group(1)
        self.assertEqual(json.loads(embedded)["runs"][0]["prompt"], payload)

    def test_review_renderer_preserves_existing_output(self):
        script = SKILLS / "review/scripts/render_review.py"
        with tempfile.TemporaryDirectory() as directory:
            source = Path(directory) / "review.json"
            output = Path(directory) / "review.html"
            source.write_text(json.dumps({"verdict": "APPROVE", "body": "</script><script>alert(1)</script>"}))
            output.write_text("user content")
            command = [sys.executable, "-B", str(script), "--review-json", str(source), "--output", str(output)]
            refused = subprocess.run(command, capture_output=True, text=True)
            self.assertNotEqual(refused.returncode, 0)
            self.assertEqual(output.read_text(), "user content")
            forced = subprocess.run(command + ["--force"], capture_output=True, text=True)
            self.assertEqual(forced.returncode, 0, forced.stderr)
            self.assertIn("\\u003c/script\\u003e", output.read_text())
            self.assertNotIn("</script><script>alert(1)</script>", output.read_text())

    def test_sentry_errors_do_not_include_sensitive_response_or_query(self):
        spec = importlib.util.spec_from_file_location("sentry_api", SKILLS / "sentry/scripts/sentry_api.py")
        sentry = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(sentry)
        url = "https://sentry.invalid/api/0/issues/?query=private@example.com"
        errors = [
            HTTPError(url, 403, "Forbidden", {}, io.BytesIO(b"secret-token private@example.com")),
            URLError("secret-token private@example.com"),
        ]
        for error in errors:
            with self.subTest(error=type(error).__name__):
                with patch.object(sentry, "urlopen", side_effect=error):
                    with self.assertRaises(RuntimeError) as raised:
                        sentry.request_json(url, "secret-token", retries=0)
                self.assertNotIn("secret-token", str(raised.exception))
                self.assertNotIn("private@example.com", str(raised.exception))


@unittest.skipUnless(importlib.util.find_spec("yaml"), "skill packaging requires PyYAML")
class SkillPackagingTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        sys.path.insert(0, str(SKILLS / "skill-creator"))
        from scripts.package_skill import package_skill
        from scripts.quick_validate import validate_skill
        cls.package = staticmethod(package_skill)
        cls.validate = staticmethod(validate_skill)

    def make_skill(self, directory):
        skill = Path(directory) / "sample"
        skill.mkdir()
        (skill / "SKILL.md").write_text("---\nname: sample\ndescription: Sample skill\n---\n# Sample\n")
        return skill

    def test_empty_required_metadata_is_invalid(self):
        with tempfile.TemporaryDirectory() as directory:
            skill = self.make_skill(directory)
            for name, description in [("''", "Valid description"), ("sample", "'   '")]:
                with self.subTest(name=name, description=description):
                    (skill / "SKILL.md").write_text(f"---\nname: {name}\ndescription: {description}\n---\n")
                    self.assertFalse(self.validate(skill)[0])

    def test_packaging_preserves_archives_and_excludes_itself(self):
        with tempfile.TemporaryDirectory() as directory:
            skill = self.make_skill(directory)
            output = skill / "dist"
            archive = self.package(skill, output)
            self.assertIsNotNone(archive)
            with zipfile.ZipFile(archive) as contents:
                self.assertEqual(contents.namelist(), ["sample/SKILL.md"])
            original = archive.read_bytes()
            self.assertIsNone(self.package(skill, output))
            self.assertEqual(archive.read_bytes(), original)

    def test_packaging_rejects_external_symlink(self):
        with tempfile.TemporaryDirectory() as directory:
            skill = self.make_skill(directory)
            private = Path(directory) / "private.txt"
            private.write_text("private")
            (skill / "leak.txt").symlink_to(private)
            self.assertIsNone(self.package(skill, Path(directory) / "output"))
            self.assertFalse((Path(directory) / "output/sample.skill").exists())


if __name__ == "__main__":
    unittest.main()
