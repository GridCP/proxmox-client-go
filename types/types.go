package types

type AuthResponse struct {
	Data struct {
		Ticket              string `json:"ticket"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Username            string `json:"username"`
	} `json:"data"`
}

type Node struct {
	Node   string `json:"node"`
	Status string `json:"status"`
}

type NodeListResponse struct {
	Data []Node `json:"data"`
}
