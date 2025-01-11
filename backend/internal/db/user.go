package db

import (
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func UserInit(id string, data *types.User) error {
	return SaveData[*types.User](BucketUsers, id, data)
}

func UserFindSupperAdmin() (*types.User, error) {
	users, err := GetDataAll[types.User](BucketUsers)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if types.IsSupperAdmin(user.Role) {
			return user, nil
		}
	}
	return nil, nil
}

func UserGetAll() ([]*types.User, error) {
	return GetDataAll[types.User](BucketUsers)
}

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
		if user.Username == name {
			if user.Password == psd {
				return user, nil
			}
			return nil, response.NewBadRequestErr(locales.CodePasswordIsWrong)
		}
	}
	return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
}

func UserUpdate(id string, data *types.User) error {
	if err := SaveData[*types.User](BucketUsers, id, data); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}
