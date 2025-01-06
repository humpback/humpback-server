package types

type Node struct {
	NodeID    string      `json:"nodeId"`
	Name      string      `json:"name"`
	IpAddress string      `json:"ipAddress"`
	Port      int         `json:"port"`
	Status    string      `json:"status"`
	CreatedAt int64       `json:"createdAt"`
	UpdatedAt int64       `json:"updatedAt"`
	NodeInfo  interface{} `json:"nodeInfo"`
}

type NodesGroups struct {
	GroupID   string   `json:"groupId"`
	GroupName string   `json:"groupName"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	Nodes     []string `json:"nodes"`
}
