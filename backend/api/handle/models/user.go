package models

import (
	"humpback/common/locales"
	"humpback/common/verify"
	"humpback/pkg/utils"
)

type UserLoginReqInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLoginReqInfo) Check() error {
	u.Username = utils.RSADecrypt(u.Username)
	u.Password = utils.RSADecrypt(u.Password)
	if err := verify.CheckIsEmpty(u.Username, locales.CodeUserNameNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckIsEmpty(u.Password, locales.CodePasswordNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckUsername(u.Username); err != nil {
		return err
	}
	if err := verify.CheckPassword(u.Password); err != nil {
		return err
	}
	return nil
}
