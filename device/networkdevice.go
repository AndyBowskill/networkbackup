package device

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type BackupNetworkDevice func(nd NetworkDevice)

type NetworkDevices struct {
	NetworkDevices []NetworkDevice `yaml:"networkdevices"`
}

type NetworkDevice struct {
	Type     string `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	IPv4     string `yaml:"ipv4"`
}

func Backup(userHomeDir, backupDir string) {

	var nds NetworkDevices

	backupNetworkCisco := func(nd NetworkDevice) {
		backupSSHToCisco(backupDir, nd.Username, nd.Password, nd.IPv4)
	}

	getConfig(userHomeDir, &nds)

	backupNetwork(&nds, backupNetworkCisco)
}

func getConfig(userHomeDir string, nds *NetworkDevices) {

	configFile, err := os.Open(userHomeDir + "/networkbackup.yaml")
	errorCheck(err)
	defer configFile.Close()

	bytes, err := io.ReadAll(configFile)
	errorCheck(err)

	yaml.Unmarshal(bytes, &nds)
}

func backupNetwork(nds *NetworkDevices, backupNetworkCisco BackupNetworkDevice) {

	for i := 0; i < len(nds.NetworkDevices); i++ {

		nd := nds.NetworkDevices[i]

		switch nd.Type {
		case cisco:
			backupNetworkCisco(nd)

		default:
			// Other vendors are not implemented yet
		}
	}
}
