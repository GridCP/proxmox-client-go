package nodes

import (
	"encoding/json"
	"gridcp-proxmox-go/client"
	"gridcp-proxmox-go/types"
)

func GetNodes(c *client.Client) ([]types.Node, error) {
	body, err := c.Get("/api2/json/nodes")
	if err != nil {
		return nil, err
	}

	var result types.NodeListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
