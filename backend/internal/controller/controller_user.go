package controller

import (
	"fmt"
	"log/slog"
	"time"

	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/config"
	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"
)

func InitAdminUser() error {
	adminConfig := config.AdminArgs()
	user, err := db.UserFindSupperAdmin()
	if err != nil {
		return fmt.Errorf("Check admin account failed: %s", err)
	}
	if user == nil {
		var (
			t  = time.Now().UnixMilli()
			id = utils.NewGuidStr()
		)
		if err = db.UserInit(id, &types.User{
			UserID:    id,
			Username:  adminConfig.Username,
			Email:     "",
			Password:  adminConfig.Password,
			Phone:     "",
			Role:      types.UserRoleSupperAdmin,
			CreatedAt: t,
			UpdatedAt: t,
			Groups:    nil,
		}); err != nil {
			return fmt.Errorf("Create admin account failed: %s", err)
		}
	}
	slog.Info("Admin account check success")
	return nil
}

func UserLogin(info *models.UserLoginReqInfo) (*types.User, string, error) {
	userInfo, err := db.UserGetByNamePsd(info.Username, info.Password)
	if err != nil {
		return nil, "", err
	}
	sessionInfo := &types.Session{
		SessionId: utils.NewGuidStr(),
		UserId:    userInfo.UserID,
	}
	if err = SessionUpdate(sessionInfo); err != nil {
		return nil, "", err
	}
	return userInfo, sessionInfo.SessionId, nil
}

func UserUpdate(userInfo *types.User) error {
	if err := db.UserUpdate(userInfo.UserID, userInfo); err != nil {
		return err
	}
	return nil
}

func UserChangePassword(userInfo *types.User, reqInfo *models.UserChangePasswordReqInfo) error {
	if userInfo.Password != reqInfo.OldPassword {
		return response.NewBadRequestErr(locales.CodeOldPasswordIsWrong)
	}
	userInfo.Password = reqInfo.NewPassword
	userInfo.UpdatedAt = time.Now().UnixMilli()
	if err := db.UserUpdate(userInfo.UserID, userInfo); err != nil {
		return err
	}
	return nil

}
