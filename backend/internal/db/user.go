package db

import (
	"strings"

	"humpback/common/response"
	"humpback/types"
)

func UserGetBySessionId(sessionId string) (*types.User, error) {
	return nil, nil
}

func UserGetById(id string) (*types.User, error) {
	return GetDataById[types.User](BucketUsers, id)
}

func UserGetByNamePsd(name string, psd string) (*types.User, error) {
	users, err := GetDataAll[types.User](BucketUsers)
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	isEmail := strings.Contains(name, "@")
	for _, user := range users {
		if isEmail {
			if user.Email == name && user.Password == psd {
				return user, nil
			}
			continue
		}
		if user.UserName == name && user.Password == psd {
			return user, nil
		}
	}
	return nil, nil
}

func UserUpdate(id string, data *types.User) error {
	return SaveData[*types.User](BucketUsers, id, data)
}
