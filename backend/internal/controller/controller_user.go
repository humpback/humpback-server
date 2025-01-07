package controller

import (
	"humpback/api/handle/models"
	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"
)

func UserLogin(info *models.UserLoginReqInfo) (*types.User, string, error) {
	userInfo, err := db.UserGetByNamePsd(info.Name, info.Password)
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
