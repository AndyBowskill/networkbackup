package device

import (
	"bufio"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

const cisco string = "cisco"

func backupSSHToCisco(username, password, ipv4 string) {

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", ipv4, sshConfig)
	if err != nil {
		log.Fatal(err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	session.Stderr = os.Stderr

	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	backupFile, err := os.Create(userHomeDir + "/cisco.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer backupFile.Close()

	writer := bufio.NewWriter(backupFile)
	defer writer.Flush()

	backupCiscoCmds(stdin, password)

	go io.Copy(writer, stdout)
	session.Wait()

}

func backupCiscoCmds(stdin io.Writer, password string) {

	stdin.Write([]byte("enable\n"))
	stdin.Write([]byte(password + "\n"))

	// Terminal length is 0 because we want to show all the options in one page
	stdin.Write([]byte("terminal length 0\n"))

	// show running-config brief
	stdin.Write([]byte("show run brief\n"))

	// Reverse the terminal length we temporarily modified
	stdin.Write([]byte("terminal no length 0\n"))

	stdin.Write([]byte("exit\n"))

}
