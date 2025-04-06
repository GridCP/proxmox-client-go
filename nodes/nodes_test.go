package nodes_test

import (
	"gridcp-proxmox-go"
	"os"
	"strconv"
	"testing"
)

func TestGetNodesSuccess(t *testing.T) {

	BaseUrl := os.Getenv("PROXMOX_URL")
	t.Logf("BaseURL: " + BaseUrl)
	UserName := os.Getenv("PROXMOX_USER")
	t.Logf("UserName: " + UserName)
	Password := os.Getenv("PROXMOX_PASSWORD")

	p, err := gridcp_proxmox_go.New(BaseUrl, UserName, Password)

	nodes, err := p.GetNodes()

	if err != nil {
		t.Fatalf("Error in obtain Nodes %v", err)
	}

	t.Logf(strconv.Itoa(len(nodes)))
	if len(nodes) == 0 {
		t.Errorf("GetNodes returned no nodes")
	}
}
