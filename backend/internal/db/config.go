package db

import (
	"strings"

	"humpback/types"
)

func ConfigsGetAll() ([]*types.Config, error) {
	return GetDataAll[types.Config](BucketConfigs)
}

func ConfigGetById(id string) (*types.Config, error) {
	return GetDataById[types.Config](BucketConfigs, id)
}

func ConfigsGetByName(name string, isLower bool) ([]*types.Config, error) {
	configs, err := GetDataAll[types.Config](BucketConfigs)
	if err != nil {
		return nil, err
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
	return SaveData[*types.Config](BucketConfigs, info.ConfigId, info)
}

func ConfigDelete(id string) error {
	return DeleteData(BucketConfigs, id)
}
