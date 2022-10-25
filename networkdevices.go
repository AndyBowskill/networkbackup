package main

type NetworkDevices struct {
	NetworkDevices []NetworkDevice `json:"networkdevices"`
}

type NetworkDevice struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	IPv4     string `json:"ipv4"`
}
