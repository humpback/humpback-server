package types

type Registry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsDefault string `json:"isDefault"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Config struct {
	ConfigID    string      `json:"configId"`
	ConfigName  string      `json:"configName"`
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
