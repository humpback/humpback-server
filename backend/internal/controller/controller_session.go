package controller

import (
	"log/slog"
	"time"

	"humpback/config"
	"humpback/internal/db"
	"humpback/types"
)

func SessionGCInterval(stopCh <-chan struct{}) {
	slog.Info("[GC Session] startup.", "Interval", config.DBArgs().SessionGCInterval.String())
	ticker := time.NewTicker(config.DBArgs().SessionGCInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			slog.Info("[GC Session] time is up...")
			sessions, err := db.SessionGetAll()
			if err != nil {
				slog.Error("[GC Session] get all session failed.", "Error", err)
			}
			var (
				nowT  = time.Now().UnixMilli()
				gcIds = make([]string, 0)
			)
			for _, session := range sessions {
				if session.ExpiredAt < nowT {
					gcIds = append(gcIds, session.SessionId)
				}
			}
			slog.Info("[GC Session] checked.", "Total", len(gcIds))
			if len(gcIds) > 0 {
				if err = db.SessionGCByIds(gcIds); err != nil {
					slog.Error("[GC Session] gc session failed.", "Total", len(gcIds), "Error", err)
				} else {
					slog.Info("[GC Session] gc session success", "Total", len(gcIds))
				}
			}

		case <-stopCh:
			return
		}
	}
}

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
