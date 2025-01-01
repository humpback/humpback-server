package models

type UserLoginReqInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserLoginReqInfo) Check() error {
	return nil
}
