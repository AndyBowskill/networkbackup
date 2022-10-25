package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {

	var nds NetworkDevices

	readJSONFile(&nds)

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

func readJSONFile(nds *NetworkDevices) {

	jsonFile, err := os.Open("networkbackup.json")
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
