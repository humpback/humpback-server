package types

type Node struct {
    NodeId      string            `json:"nodeId"`
    Name        string            `json:"name"`
    IpAddress   string            `json:"ipAddress"`
    Port        int               `json:"port"`
    Status      string            `json:"status"`
    IsEnable    bool              `json:"isEnable"`
    CreatedAt   int64             `json:"createdAt"`
    UpdatedAt   int64             `json:"updatedAt"`
    CPUUsage    float32           `json:"cpuUsage"`
    MemoryUsage float32           `json:"memoryUsage"`
    Labels      map[string]string `json:"labels"`
}

type NodesGroups struct {
    GroupId   string   `json:"groupId"`
    GroupName string   `json:"groupName"`
    CreatedAt int64    `json:"createdAt"`
    UpdatedAt int64    `json:"updatedAt"`
    Nodes     []string `json:"nodes"`
}
