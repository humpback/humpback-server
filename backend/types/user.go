package types

type UserRole int

var (
	UserRoleSupperAdmin UserRole = 1
	UserRoleAdmin       UserRole = 2
	UserRoleNormal      UserRole = 3
)

type User struct {
	UserID      string   `json:"userId"`
	Username    string   `json:"username"`
	Description string   `json:"description"`
	Email       string   `json:"email"`
	Password    string   `json:"password,omitempty"`
	Phone       string   `json:"phone"`
	Role        UserRole `json:"role"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   int64    `json:"updatedAt"`
	Groups      []string `json:"groups"`
}

type Group struct {
	GroupID     string   `json:"groupId"`
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

func IsNormal(role UserRole) bool {
	return role == UserRoleNormal
}
