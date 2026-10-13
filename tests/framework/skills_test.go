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
