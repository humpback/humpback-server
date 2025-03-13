package types

type GroupContainerPerformance struct {
	ContainerId string               `json:"containerId"`
	NodeId      string               `json:"nodeId"`
	IsSuccess   bool                 `json:"isSuccess"`
	Error       string               `json:"error"`
	Stats       *GroupContainerStats `json:"stats"`
}

type GroupContainerStats struct {
	CpuPercent  float64                   `json:"cpuPercent"`
	MermoryUsed uint64                    `json:"memoryUsed"`
	MemoryLimit uint64                    `json:"memoryLimit"`
	IORead      uint64                    `json:"ioRead"`
	IOWrite     uint64                    `json:"ioWrite"`
	StatsAt     int64                     `json:"statsAt"`
	Networks    []*ContainerNetWorkStatus `json:"networks"`
}

type ContainerNetWorkStatus struct {
	Name    string `json:"name"`
	RxBytes uint64 `json:"rxBytes"`
	TxBytes uint64 `json:"txBytes"`
}
