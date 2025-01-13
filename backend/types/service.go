package types

const (
	NodeStatusOnline  = "Online"
	NodeStatusOffline = "Offline"
)

const (
	ServiceStatusNotReady = "NotReady"
	ServiceStatusRunning  = "Running"
	ServiceStatusFailed   = "Failed"

	ContainerStatusPending  = "Pending"
	ContainerStatusStarting = "Starting"
	ContainerStatusCreated  = "Created"
	ContainerStatusRunning  = "Running"
	ContainerStatusFailed   = "Failed"
	ContainerStatusExited   = "Exited"
	ContainerStatusRemoved  = "Removed"
	ContainerStatusWarning  = "Warning"
)

type Service struct {
	ServiceId   string             `json:"serviceId"`
	ServiceName string             `json:"serviceName"`
	Version     string             `json:"version"`
	IsEnabled   bool               `json:"isEnabled"`
	Status      string             `json:"status"`
	Meta        *ServiceMetaDocker `json:"meta"`
	Deployment  *Deployment        `json:"deployment"`
	Containers  []*ContainerStatus `json:"containers"`
	GroupId     string             `json:"groupId"`
	CreateAt    int64              `json:"createAt"`
	UpdateAt    int64              `json:"updateAt"`
}

type ContainerStatus struct {
	ContainerId   string `json:"containerId"`
	ContainerName string `json:"containerName"`
	NodeId        string `json:"nodeId"`
	Status        string `json:"state"`
	ErrorMsg      string `json:"errorMsg"`
	Image         string `json:"image"`
	Command       string `json:"command"`
	CreateAt      int64  `json:"createAt"`
	StartAt       int64  `json:"startAt"`
}

type Deployment struct {
	Type       DeployType       `json:"type"`
	Mode       DeployMode       `json:"mode"`
	Replicas   int              `json:"replicas"`
	Placements []*PlacementInfo `json:"placements"`
	Schedule   *ScheduleInfo    `json:"schedule"`
}

type DeployMode string
type DeployType string

var (
	DeployModeGlobal    DeployMode = "global"
	DeployModeReplicate DeployMode = "replicate"
)

var (
	DeployTypeSchedule   DeployType = "schedule"
	DeployTypeBackground DeployType = "background"
)

type PlacementMode string

var (
	PlacementModeLabel PlacementMode = "label"
	PlacementModeIP    PlacementMode = "ip"
)

type PlacementInfo struct {
	Mode    PlacementMode `json:"mode"`
	Key     string        `json:"key"`
	Value   string        `json:"value"`
	IsEqual bool          `json:"isEqual"`
}

// Rule is cron string
type ScheduleInfo struct {
	Timeout string   `json:"timeout"`
	Rules   []string `json:"rules"`
}

// type ScheduleRule struct {
// 	Id          string `json:"id"`
// 	Enabled     bool   `json:"enabled"`
// 	Mode        string `json:"mode"`
// 	StartDateAt int64  `json:"startAt"`
// 	EndDateAt   int64  `json:"endAt"`
// 	StartTimeAt int64  `json:"startTimeAt"`
// }

type ServiceMetaDocker struct {
	Image         string            `json:"image"`
	AlwaysPull    bool              `json:"alwaysPull"`
	Command       string            `json:"command"`
	Envs          []string          `json:"env"`
	Labels        map[string]string `json:"labels"`
	Network       *NetworkInfo      `json:"network"`
	RestartPolicy *RestartPolicy    `json:"restartPolicy"`
}

type NetworkMode string

var (
	NetworkModeHost   NetworkMode = "host"
	NetworkModeBridge NetworkMode = "bridge"
	NetworkModeCustom NetworkMode = "custom"
)

type NetworkInfo struct {
	Mode        NetworkMode `json:"mode"`        // custom模式需要创建网络
	Hostname    string      `json:"hostname"`    // bridge及custom模式时可设置，用户容器的hostname
	NetworkName string      `json:"networkName"` //custom模式使用
	Ports       []*PortInfo `json:"ports"`
}

type PortInfo struct {
	HostPort      uint   `json:"hostPort"`
	ContainerPort uint   `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

type RestartPolicyMode string

var (
	RestartPolicyModeNo            RestartPolicyMode = "no"
	RestartPolicyModeAlways        RestartPolicyMode = "always"
	RestartPolicyModeOnFail        RestartPolicyMode = "on-failure"
	RestartPolicyModeUnlessStopped RestartPolicyMode = "unless-stopped"
)

type RestartPolicy struct {
	Mode          RestartPolicyMode `json:"mode"`
	MaxRetryCount int               `json:"maxRetryCount"`
}
