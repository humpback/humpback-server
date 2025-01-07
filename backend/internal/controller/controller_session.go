package controller

import (
	"time"

	"humpback/config"
	"humpback/internal/db"
	"humpback/types"
)

func SessionGetAndRefresh(sessionId string) (*types.User, bool, error) {
	sessionInfo, expired, err := db.SessionGetById(sessionId)
	if err != nil {
		return nil, expired, err
	}
	if expired {
		return nil, true, nil
	}
	userInfo, err := db.UserGetById(sessionInfo.UserId)
	if err != nil {
		return nil, false, err
	}
	userInfo.Password = ""
	if err = SessionUpdate(sessionInfo); err != nil {
		return nil, false, err
	}
	return userInfo, false, nil
}

func SessionUpdate(sessionInfo *types.Session) error {
	sessionInfo.ExpiredAt = time.Now().Add(config.DBArgs().SessionTimeout).UnixMilli()
	return db.SessionUpdate(sessionInfo)
}

func SessionDelete(sessionId string) error {
	return db.SessionDelete(sessionId)
}
