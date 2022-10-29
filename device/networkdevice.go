package device

import (
	"encoding/json"
	"io"
	"log"
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

func Backup() {

	var nds NetworkDevices

	backupNetworkCisco := func(nd NetworkDevice) {
		backupSSHToCisco(nd.Username, nd.Password, nd.IPv4)
	}

	getConfig(&nds, "networkbackup.json")

	backupNetwork(&nds, backupNetworkCisco)

}

func getConfig(nds *NetworkDevices, backupFile string) {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(userHomeDir + "/" + backupFile)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

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
