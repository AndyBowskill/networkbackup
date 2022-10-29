package device

import (
	"encoding/json"
	"testing"
)

func TestBackupNetworkValidArgs(t *testing.T) {

	var nds NetworkDevices

	bytes := []byte("\n{\n\t\"networkdevices\": [{\n\t\t\"type\": \"cisco\",\n\t\t\"username\": \"andybowskill\",\n\t\t\"password\": \"cisco\",\n\t\t\"ipv4\": \"192.168.48.2:22\"\n\t}]\n}")
	json.Unmarshal(bytes, &nds)

	backupToCisco := func(nd NetworkDevice) {}

	backupNetwork(&nds, backupToCisco)

	got := len(nds.NetworkDevices)
	want := 1

	if got != want {
		t.Errorf("backupNetworkDevice function didn't contain one NetworkDevice element only")
	}
}

func TestBackupNetworkInvalidArgs(t *testing.T) {

	var nds NetworkDevices

	bytes := []byte("\n{\n\t\"networkdevices\": [")
	json.Unmarshal(bytes, &nds)

	backupToCisco := func(nd NetworkDevice) {}

	backupNetwork(&nds, backupToCisco)

	got := len(nds.NetworkDevices)
	want := 0

	if got != want {
		t.Errorf("backupNetworkDevice function didn't contain zero NetworkDevice element only")
	}
}
