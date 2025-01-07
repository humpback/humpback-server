package db

import (
	"time"

	"humpback/common/response"
	"humpback/types"
)

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
		return response.NewBadRequestErr(err.Error())
	}
	return nil
}

func SessionDelete(sessionId string) error {
	if err := DeleteData(BucketSessions, sessionId); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}
