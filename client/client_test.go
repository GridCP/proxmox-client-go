package client_test

import (
	"gridcp-proxmox-go"
	"os"
	"testing"
)

func TestLoginSuccess(t *testing.T) {

	BaseUrl := os.Getenv("PROXMOX_URL")
	t.Logf("BaseURL: " + BaseUrl)
	UserName := os.Getenv("PROXMOX_USER")
	t.Logf("UserName: " + UserName)
	Password := os.Getenv("PROXMOX_PASSWORD")

	p, err := gridcp_proxmox_go.New(BaseUrl, UserName, Password)

	if err != nil {
		t.Fatalf("Auth Error %v", err)
	}

	t.Logf("Connect to Proxmox")

	nodes, err := p.GetNodes()
	if err != nil {
		t.Fatalf("Node fetch error: %v", err)
	}

	for _, node := range nodes {
		t.Logf("Node: %s, Status: %s\n", node.Node, node.Status)
	}

	if len(nodes) == 0 {
		t.Error("Expected at least one node, got none")
	} else {
		t.Logf("Connected successfully. Found %d nodes", len(nodes))
	}
}
