package device

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type NetworkDevices struct {
	NetworkDevices []NetworkDevice `json:"networkdevices"`
}

type NetworkDevice struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	IPv4     string `json:"ipv4"`
}

func BackupNetwork() {

	var nds NetworkDevices

	readNetworkDeviceFile(&nds)

	for i := 0; i < len(nds.NetworkDevices); i++ {

		nd := nds.NetworkDevices[i]

		switch nd.Type {
		case cisco:
			sshToCisco(nd.Username, nd.Password, nd.IPv4)

		default:
			// Other vendors are not implemented yet
		}
	}

}

func readNetworkDeviceFile(nds *NetworkDevices) {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(userHomeDir + "/networkbackup.json")
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
