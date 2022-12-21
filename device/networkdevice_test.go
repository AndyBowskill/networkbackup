package device

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func Test_backupNetworkSupportedVendor(t *testing.T) {

	bytes := []byte("networkdevices:\r\n  - type: cisco\r\n    username: andybowskill\r\n    password: cisco\r\n    ipv4: '192.168.48.2:22'\r\n")

	backupNetworkArgs(t, bytes, 1)
}

func Test_backupNetworkUnsupportedVendor(t *testing.T) {

	bytes := []byte("networkdevices:\r\n  - type: aruba\r\n    username: andybowskill\r\n    password: pass123\r\n    ipv4: '192.168.48.2:22'\r\n")

	backupNetworkArgs(t, bytes, 1)
}

func Test_backupNetworkInvalidArgs(t *testing.T) {

	bytes := []byte("networkdevices:")

	backupNetworkArgs(t, bytes, 0)
}

func backupNetworkArgs(t *testing.T, bytes []byte, want int) {

	var nds NetworkDevices

	yaml.Unmarshal(bytes, &nds)

	backupToCisco := func(nd NetworkDevice) {}

	backupNetwork(&nds, backupToCisco)

	got := len(nds.NetworkDevices)

	if got != want {
		t.Errorf("Expected NetworkDevice element of %d isn't the same actual of %d", got, want)
	}

}
