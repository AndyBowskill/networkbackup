package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {

	sshConfig := &ssh.ClientConfig{
		User: "andybowskill",
		Auth: []ssh.AuthMethod{ssh.Password("cisco")},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", "192.168.48.2:22", sshConfig)
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
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	stdin.Write([]byte("enable\n"))
	stdin.Write([]byte("cisco\n"))

	// Terminal length is 0 because we want to show all the options in one page
	stdin.Write([]byte("terminal length 0\n"))

	stdin.Write([]byte("show run brief\n"))
	time.Sleep(1 * time.Second)

	// Reverse the terminal length we temporarily modified
	stdin.Write([]byte("terminal no length 0\n"))

}
