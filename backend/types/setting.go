package types

type Registry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsDefault string `json:"isDefault"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type ConfigType int

var (
	ConfigTypeStatic ConfigType = 1
	ConfigTypeVolume ConfigType = 2
)

type Config struct {
	ConfigId    string      `json:"configId"`
	ConfigName  string      `json:"configName"`
	Description string      `json:"description"`
	ConfigType  ConfigType  `json:"configType"`
	ConfigValue interface{} `json:"configValue"`
	CreateAt    int64       `json:"createAt"`
	UpdateAt    int64       `json:"updateAt"`
}

type Template struct {
	TemplateID   string      `json:"templateId"`
	TemplateName string      `json:"templateName"`
	Content      interface{} `json:"content"`
	CreateAt     int64       `json:"createAt"`
	UpdateAt     int64       `json:"updateAt"`
}
