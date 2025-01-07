package models

import (
	"strings"

	"humpback/common/verify"
	"humpback/pkg/utils"
)

type UserLoginReqInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserLoginReqInfo) Check() error {
	u.Name = utils.RSADecrypt(u.Name)
	u.Password = utils.RSADecrypt(u.Password)
	if strings.Contains(u.Name, "@") {
		if err := verify.CheckEmail(u.Name); err != nil {
			return err
		}
	} else {
		if err := verify.CheckName(u.Name); err != nil {
			return err
		}
	}
	if err := verify.CheckPassword(u.Password); err != nil {
		return err
	}
	return nil
}
