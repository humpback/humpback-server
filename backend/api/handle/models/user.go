package models

import (
	"encoding/json"
	"time"

	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/pkg/utils"
	"humpback/types"
)

type UserLoginReqInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLoginReqInfo) Check() error {
	u.Username = utils.RSADecrypt(u.Username)
	u.Password = utils.RSADecrypt(u.Password)

	if err := verify.CheckRequiredAndLengthLimit(u.Username, locales.LimitUserName.Min, locales.LimitUserName.Max, locales.CodeUserNameNotEmpty, locales.CodeUserNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckRequiredAndLengthLimit(u.Password, locales.LimitPassword.Min, locales.LimitPassword.Max, locales.CodePasswordNotEmpty, locales.CodePasswordLimitLength); err != nil {
		return err
	}
	return nil
}

type MeUpdateReqInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

func (u *MeUpdateReqInfo) Check() error {
	if err := verify.CheckRequiredAndLengthLimit(u.Username, locales.LimitUserName.Min, locales.LimitUserName.Max, locales.CodeUserNameNotEmpty, locales.CodeUserNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(u.Email, 0, locales.LimitEmail.Max, locales.CodeEmailLimitLength); err != nil {
		return err
	}
	if u.Email != "" {
		if err := verify.CheckEmail(u.Email); err != nil {
			return err
		}
	}
	if err := verify.CheckLengthLimit(u.Phone, 0, locales.LimitPhone.Max, locales.CodePhoneLimitLength); err != nil {
		return err
	}
	if u.Phone != "" {
		if err := verify.CheckPhone(u.Phone); err != nil {
			return err
		}
	}
	if err := verify.CheckLengthLimit(u.Description, 0, locales.LimitDescription.Max, locales.CodeDescriptionLimitLength); err != nil {
		return err
	}
	return nil
}

func (u *MeUpdateReqInfo) NewUserInfo(userInfo *types.User) *types.User {
	userInfo.Username = u.Username
	userInfo.Email = u.Email
	userInfo.Phone = u.Phone
	userInfo.Description = u.Description
	userInfo.UpdatedAt = time.Now().UnixMilli()
	return userInfo
}

type MeChangePasswordReqInfo struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (u *MeChangePasswordReqInfo) Check() error {
	u.OldPassword = utils.RSADecrypt(u.OldPassword)
	u.NewPassword = utils.RSADecrypt(u.NewPassword)
	if err := verify.CheckRequiredAndLengthLimit(u.OldPassword, locales.LimitPassword.Min, locales.LimitPassword.Max, locales.CodeOldPasswordNotEmpty, locales.CodeOldPasswordLimitLength); err != nil {
		return err
	}
	if err := verify.CheckRequiredAndLengthLimit(u.NewPassword, locales.LimitPassword.Min, locales.LimitPassword.Max, locales.CodeNewPasswordNotEmpty, locales.CodeNewPasswordLimitLength); err != nil {
		return err
	}
	return nil
}

type UserCreateReqInfo struct {
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Role        types.UserRole `json:"role"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	Description string         `json:"description"`
	Teams       []string       `json:"teams"`
}

func (u *UserCreateReqInfo) Check() error {
	u.Password = utils.RSADecrypt(u.Password)

	if err := verify.CheckRequiredAndLengthLimit(u.Username, locales.LimitUserName.Min, locales.LimitUserName.Max, locales.CodeUserNameNotEmpty, locales.CodeUserNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(u.Email, 0, locales.LimitEmail.Max, locales.CodeEmailLimitLength); err != nil {
		return err
	}
	if u.Email != "" {
		if err := verify.CheckEmail(u.Email); err != nil {
			return err
		}
	}
	if err := verify.CheckLengthLimit(u.Phone, 0, locales.LimitPhone.Max, locales.CodePhoneLimitLength); err != nil {
		return err
	}
	if u.Phone != "" {
		if err := verify.CheckPhone(u.Phone); err != nil {
			return err
		}
	}
	if err := verify.CheckLengthLimit(u.Description, 0, locales.LimitDescription.Max, locales.CodeDescriptionLimitLength); err != nil {
		return err
	}

	if err := verify.CheckRequiredAndLengthLimit(u.Password, locales.LimitPassword.Min, locales.LimitPassword.Max, locales.CodePasswordNotEmpty, locales.CodePasswordLimitLength); err != nil {
		return err
	}
	if u.Role != types.UserRoleAdmin && u.Role != types.UserRoleUser {
		return response.NewBadRequestErr(locales.CodeUserRoleIsInvalid)
	}
	if len(u.Teams) == 0 {
		u.Teams = make([]string, 0)
	}
	return nil
}

func (u *UserCreateReqInfo) CheckRole(role types.UserRole) error {
	if role == types.UserRoleAdmin && u.Role != types.UserRoleUser {
		return response.NewBadRequestErr(locales.CodeUserRoleIsInvalid)
	}
	return nil
}

type UserUpdateReqInfo struct {
	UserId string `json:"userId"`
	UserCreateReqInfo
}

func (u *UserUpdateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(u.UserId, locales.CodeUserIdNotEmpty); err != nil {
		return err
	}
	return u.UserCreateReqInfo.Check()
}

type UserQueryFilterInfo struct {
	Role int `json:"role"`
}

type UserQueryReqInfo struct {
	types.QueryInfo
	FilterInfo *UserQueryFilterInfo `json:"-"`
}

func (u *UserQueryReqInfo) Check() error {
	u.CheckBase()
	if err := u.parseFilter(); err != nil {
		return err
	}
	return nil
}

func (u *UserQueryReqInfo) parseFilter() error {
	if len(u.Filter) == 0 {
		return nil
	}
	v, err := json.Marshal(u.Filter)
	if err != nil {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	u.FilterInfo = new(UserQueryFilterInfo)
	if err = json.Unmarshal(v, u.FilterInfo); err != nil {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}
