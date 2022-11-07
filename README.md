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
1. The backup file will be in your networkbackup directory after the tool has been run.
	1. For example, on a macOS is ```Root``` > ```Users``` > ```johnsmith``` > ```networkbackup``` > ```cisco-2022-November-7-15-55```. The tool checks if the ```networkbackup``` directory is there. If not, the tool creates the directory automatically and place the backup files inside.
1. The tool only deals with Cisco devices. Other vendors are not implemented yet. If you are interested in this tool, please contribute!