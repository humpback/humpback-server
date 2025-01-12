package types

type Node struct {
	NodeId      string  `json:"nodeId"`
	Name        string  `json:"name"`
	IpAddress   string  `json:"ipAddress"`
	Port        int     `json:"port"`
	Status      string  `json:"status"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
	CPUUsage    float32 `json:"cpuUsage"`
	MemoryUsage float32 `json:"memoryUsage"`
}

type NodesGroups struct {
	GroupId   string   `json:"groupId"`
	GroupName string   `json:"groupName"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	Nodes     []string `json:"nodes"`
}
