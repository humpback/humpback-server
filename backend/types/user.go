package types

type UserRole int

var (
	UserRoleSupperAdmin UserRole = 1
	UserRoleAdmin       UserRole = 2
	UserRoleUser        UserRole = 3
)

type User struct {
	UserId      string   `json:"userId"`
	Username    string   `json:"username"`
	Description string   `json:"description"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Phone       string   `json:"phone"`
	Role        UserRole `json:"role"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   int64    `json:"updatedAt"`
	Teams       []string `json:"teams"`
}

type Team struct {
	TeamId      string   `json:"teamId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   int64    `json:"updatedAt"`
	Users       []string `json:"users"`
}

func IsSupperAdmin(role UserRole) bool {
	return role == UserRoleSupperAdmin
}

func IsAdmin(role UserRole) bool {
	return role == UserRoleAdmin
}

func IsUser(role UserRole) bool {
	return role == UserRoleUser
}
