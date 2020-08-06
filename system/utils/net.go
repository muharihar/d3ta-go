package utils

import (
	"net"
	"os"
)

// GetCurrentIP get current IP Address from this host
func GetCurrentIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip4 := localAddr.IP.To4()
	if ip4 == nil {
		return "127.0.0.1"
	}
	return ip4.String()
}

// GetHostName get current hostname from this host
func GetHostName() string {
	hn, err := os.Hostname()
	if err != nil {
		return "localhost"
	}
	return hn
}
