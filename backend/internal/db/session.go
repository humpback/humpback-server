package db

import (
	"time"

	"humpback/common/response"
	"humpback/types"
)

func SessionGetAll() ([]*types.Session, error) {
	return GetDataAll[types.Session](BucketSessions)
}

func SessionGCByIds(ids []string) error {
	return BatchDelete(BucketSessions, ids)
}

func SessionGetById(sessionId string) (*types.Session, bool, error) {
	sessionInfo, err := GetDataById[types.Session](BucketSessions, sessionId)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, true, nil
		}
		return nil, false, response.NewRespServerErr(err.Error())
	}
	return sessionInfo, sessionInfo.ExpiredAt < time.Now().UnixMilli(), nil
}

func SessionUpdate(data *types.Session) error {
	if err := SaveData[*types.Session](BucketSessions, data.SessionId, data); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func SessionDelete(sessionId string) error {
	if err := DeleteData(BucketSessions, sessionId); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func SessionBatchDeleteBuUserId(userId string) error {
	sessions, err := SessionGetAll()
	if err != nil {
		return response.NewRespServerErr(err.Error())
	}
	var ids []string
	for _, session := range sessions {
		if session.UserId == userId {
			ids = append(ids, session.SessionId)
		}
	}
	if len(ids) > 0 {
		if err = BatchDelete(BucketSessions, ids); err != nil {
			return response.NewRespServerErr(err.Error())
		}
	}
	return nil
}
