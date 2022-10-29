## Network Backup

[![Go Build and Test](https://github.com/andybowskill/networkbackup/actions/workflows/go.yml/badge.svg)](https://github.com/andybowskill/networkbackup/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/andybowskill/networkbackup)](https://goreportcard.com/report/github.com/andybowskill/networkbackup)

A command line tool to backup your network devices into a backup text file.

### Install

```
$ go install github.com/AndyBowskill/networkbackup@latest
```

### Example 

An example ```networkbackup.json``` file in your home directory:

```
{
	"networkdevices": [{
		"type": "cisco",
		"username": "johnsmith",
		"password": "password123",
		"ipv4": "192.168.0.2:22"
	}]
}
```

At the moment:
1. Your networkbackup.json file should be in your home directory.
    1. For example, on a macOS is ```Root``` > ```Users``` > ```johnsmith``` directory.
1. The ```cisco.txt``` backup file will be in your home directory after the tool has been run as well.