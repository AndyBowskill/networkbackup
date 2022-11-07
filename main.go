package main

import (
	"log"
	"os"

	"github.com/AndyBowskill/networkbackup/device"
)

func main() {

	userHomeDir, err := os.UserHomeDir()
	errorCheck(err)

	backupDir := userHomeDir + "/networkbackup"

	err = os.Chdir(backupDir)
	if err != nil {
		err = os.Mkdir(backupDir, 0755)
		errorCheck(err)
	}

	device.Backup(backupDir)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
