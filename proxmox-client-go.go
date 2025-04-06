package gridcp_proxmox_go

import (
	"gridcp-proxmox-go/client"
	"gridcp-proxmox-go/nodes"
	"gridcp-proxmox-go/types"
)

type Proxmox struct {
	Client *client.Client
}

func New(baseURL, username, password string) (*Proxmox, error) {
	c := client.NewClient(baseURL, username, password)
	if err := c.Login(); err != nil {
		return nil, err
	}
	return &Proxmox{Client: c}, nil
}

func (p *Proxmox) GetNodes() ([]types.Node, error) {
	return nodes.GetNodes(p.Client)
}
