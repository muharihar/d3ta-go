package utils

import "testing"

func TestNet(t *testing.T){
	ip := GetCurrentIP()
	if ip == "" {
		t.Error("Invalid IP")
	}
	t.Logf("GetCurrentIP: %v", ip)
	
	hn := GetHostName()
	if hn == "" {
		t.Error("Invalid HostName")
	}
	t.Logf("GetHostName: %v", hn)
}