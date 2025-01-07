package models

import (
	"humpback/common/locales"
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
	if err := verify.CheckIsEmpty(u.Name, locales.CodeUserNameNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckIsEmpty(u.Password, locales.CodePasswordNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckName(u.Name); err != nil {
		return err
	}
	if err := verify.CheckPassword(u.Password); err != nil {
		return err
	}
	return nil
}
