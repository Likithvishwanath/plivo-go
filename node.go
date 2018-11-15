package plivo


type Node struct {
	APIID     string `json:"api_id" url:"api_id"`
	NodeID     string `json:"node_id" url:"node_id"`
	PhloID string `json:"phlo_id" url:"phlo_id"`
	Name string `json:"name" url:"name"`
	NodeType string `json:"node_type" url:"node_type"`
	CreatedOn string `json:"created_on" url:"created_on"`
}
