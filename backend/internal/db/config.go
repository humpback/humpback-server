package db

import (
	"strings"

	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func ConfigGetAll() ([]*types.Config, error) {
	return GetDataAll[types.Config](BucketConfigs)
}

func ConfigGetById(id string) (*types.Config, error) {
	info, err := GetDataById[types.Config](BucketConfigs, id)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeConfigNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}

func ConfigsGetByName(name string, isLower bool) ([]*types.Config, error) {
	configs, err := GetDataAll[types.Config](BucketConfigs)
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	var result []*types.Config
	for _, config := range configs {
		if isLower && strings.ToLower(config.ConfigName) == strings.ToLower(name) {
			result = append(result, config)
		}
		if !isLower && config.ConfigName == name {
			result = append(result, config)
		}
	}
	return result, nil
}

func ConfigUpdate(info *types.Config) error {
	if err := SaveData[*types.Config](BucketConfigs, info.ConfigId, info); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func ConfigDelete(id string) error {
	if err := DeleteData(BucketConfigs, id); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}
