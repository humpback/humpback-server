package types

type Service struct {
	ServiceId   string      `json:"serviceId"`
	ServiceName string      `json:"serviceName"`
	Type        string      `json:"type"`
	Status      string      `json:"status"`
	Meta        interface{} `json:"meta"`
	Containers  []string    `json:"containers"`
	GroupId     string      `json:"groupId"`
	CreateAt    int64       `json:"createAt"`
	UpdateAt    int64       `json:"updateAt"`
}

type ContainerStatus struct {
	ContainerId   string `json:"containerId"`
	ContainerName string `json:"containerName"`
	Status        string `json:"status"`
	Image         string `json:"image"`
	Command       string `json:"command"`
	CreateAt      int64  `json:"createAt"`
	StartAt       int64  `json:"startAt"`
}

const ServiceStatusNotReady = "NotReady"
const ServiceStatusRunning = "Running"
const ServiceStatusFailed = "Failed"

const ContainerStatusPending = "Pending"
const ContainerStatusStarting = "Starting"
const ContainerStatusCreated = "Created"
const ContainerStatusRunning = "Running"
const ContainerStatusFailed = "Failed"
const ContainerStatusExited = "Exited"
const ContainerStatusRemoved = "Removed"
const ContainerStatusWarning = "Warning"
