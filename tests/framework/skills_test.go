package framework_test

import (
	"os/exec"
	"testing"
)

func TestWalkthroughScripts(t *testing.T) {
	command := exec.Command("python3", "-B", "walkthrough_test.py")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("walkthrough scripts: %v\n%s", err, output)
	}
	t.Log(string(output))
}

func TestSecurityAuditScripts(t *testing.T) {
	command := exec.Command("node", "--test",
		"security-audit/validate-findings.test.cjs",
		"security-audit/validate-coverage-ledger.test.cjs")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("security audit scripts: %v\n%s", err, output)
	}
	t.Log(string(output))
}

func TestBeforeAndAfterScript(t *testing.T) {
	command := exec.Command("node", "--test", "before-and-after.test.mjs")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("before-and-after script: %v\n%s", err, output)
	}
	t.Log(string(output))
}

func TestSkillHelpers(t *testing.T) {
	command := exec.Command("python3", "-B", "skill_helpers_test.py")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("skill helpers: %v\n%s", err, output)
	}
	t.Log(string(output))
}
