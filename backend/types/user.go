package types

type User struct {
	UserID    string   `json:"userId"`
	UserName  string   `json:"userName"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Phone     string   `json:"phone"`
	IsAdmin   bool     `json:"isAdmin"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
	Groups    []string `json:"groups"`
}

type Group struct {
	GroupID     string   `json:"groupid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   int64    `json:"updatedAt"`
	Users       []string `json:"users"`
}
