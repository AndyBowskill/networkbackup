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
	password := "12345"
	showrunCmd := "show run brief"

	backupCiscoCmds(in, password)

	got := b.String()

	if !strings.Contains(got, password) {
		t.Errorf("backupCisco function didn't contain a valid password of %s.", password)
	}
	if !strings.Contains(got, showrunCmd) {
		t.Errorf("backupCisco function didn't contain a valid show run command of %s.", showrunCmd)
	}
}
