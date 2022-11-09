package device

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestBackupCiscoCmds(t *testing.T) {

	var b = new(bytes.Buffer)
	var in io.Writer = b

	enable := "enable"
	password := "12345"
	termLength := "terminal length 0"
	showrunCmd := "show run brief"
	noTermLength := "terminal no length 0"
	exit := "exit"

	backupCiscoCmds(in, password)

	got := b.String()

	if !strings.Contains(got, enable) {
		t.Errorf("Expected Cisco device contains '%s'", enable)
	}
	if !strings.Contains(got, password) {
		t.Errorf("Expected device password of '%s'", password)
	}
	if !strings.Contains(got, termLength) {
		t.Errorf("Expected Cisco device contains '%s'", termLength)
	}
	if !strings.Contains(got, showrunCmd) {
		t.Errorf("Expected show run command of '%s'", showrunCmd)
	}
	if !strings.Contains(got, noTermLength) {
		t.Errorf("Expected Cisco device contains '%s'", noTermLength)
	}
	if !strings.Contains(got, exit) {
		t.Errorf("Expected Cisco device contains '%s'", exit)
	}

}
