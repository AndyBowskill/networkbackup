package device

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBackupNetworkValidArgs(t *testing.T) {

	var nds NetworkDevices

	bytes := []byte("networkdevices:\r\n  - type: cisco\r\n    username: andybowskill\r\n    password: cisco\r\n    ipv4: '192.168.48.2:22'\r\n")
	yaml.Unmarshal(bytes, &nds)

	backupToCisco := func(nd NetworkDevice) {}

	backupNetwork(&nds, backupToCisco)

	got := len(nds.NetworkDevices)
	want := 1

	if got != want {
		t.Errorf("Expected NetworkDevice element of %d isn't the same actual of %d", got, want)
	}
}

func TestBackupNetworkInvalidArgs(t *testing.T) {

	var nds NetworkDevices

	bytes := []byte("networkdevices:")
	yaml.Unmarshal(bytes, &nds)

	backupToCisco := func(nd NetworkDevice) {}

	backupNetwork(&nds, backupToCisco)

	got := len(nds.NetworkDevices)
	want := 0

	if got != want {
		t.Errorf("Expected NetworkDevice element of %d isn't the same actual of %d", got, want)
	}
}
