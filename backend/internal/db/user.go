package db

import (
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func UserGetById(id string) (*types.User, error) {
	info, err := GetDataById[types.User](BucketUsers, id)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}

func UserGetByNamePsd(name string, psd string) (*types.User, error) {
	users, err := GetDataAll[types.User](BucketUsers)
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	for _, user := range users {
		if user.UserName == name {
			if user.Password == psd {
				return user, nil
			}
			return nil, response.NewBadRequestErr(locales.CodePasswordIsWrong)
		}
	}
	return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
}

func UserUpdate(id string, data *types.User) error {
	return SaveData[*types.User](BucketUsers, id, data)
}
