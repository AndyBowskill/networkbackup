package device

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type FileOpen func(filename string) (*os.File, error)
type FileRead func(filename string) ([]byte, error)

type ConfigRepository struct {
	fileOpen FileOpen
	fileRead FileRead
}

func NewConfigRepository(fileOpen FileOpen, fileRead FileRead) ConfigRepository {
	return ConfigRepository{
		fileRead: fileRead,
	}
}

func (r *ConfigRepository) GetConfiguration(nds *NetworkDevices, backupFile string) {

	jsonFile, err := r.fileOpen("networkbackup.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	bytes, err := r.fileRead(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(bytes, &nds)
}

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

	var fileRead FileRead
	var nds NetworkDevices

	fileOpen = func(filename string) (*os.File, error) {

		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		jsonFile, err := os.Open(userHomeDir + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		return jsonFile, nil
	}

	fileRead = func(jsonFile string) ([]byte, error) {

		bytes, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err)
		}

		return bytes, nil
	}

	repository := NewConfigurationRepository(fileRead)

	repository.GetConfiguration(&nds, "networkbackup.json")

	for i := 0; i < len(nds.NetworkDevices); i++ {

		nd := nds.NetworkDevices[i]

		switch nd.Type {
		case cisco:
			backupSSHToCisco(nd.Username, nd.Password, nd.IPv4)

		default:
			// Other vendors are not implemented yet
		}
	}

}

/*
func readNetworkDeviceFile(nds *NetworkDevices, backupFile string) {

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

}*/
