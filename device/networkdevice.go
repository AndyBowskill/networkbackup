package device

import (
	"encoding/json"
	"io"
	"os"
)

type BackupNetworkDevice func(nd NetworkDevice)

type NetworkDevices struct {
	NetworkDevices []NetworkDevice `json:"networkdevices"`
}

type NetworkDevice struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	IPv4     string `json:"ipv4"`
}

func Backup(backupDir string) {

	var nds NetworkDevices

	backupNetworkCisco := func(nd NetworkDevice) {
		backupSSHToCisco(backupDir, nd.Username, nd.Password, nd.IPv4)
	}

	getConfig(&nds, backupDir)

	backupNetwork(&nds, backupNetworkCisco)

}

func getConfig(nds *NetworkDevices, backupDir string) {

	jsonFile, err := os.Open(backupDir + "/networkbackup.json")
	errorCheck(err)

	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	errorCheck(err)

	json.Unmarshal(bytes, &nds)
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
