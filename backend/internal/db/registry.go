package db

import (
	"encoding/json"
	"fmt"
	"strings"

	bolt "go.etcd.io/bbolt"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func RegistryGetAll() ([]*types.Registry, error) {
	return GetDataAll[types.Registry](BucketRegistries)
}

func RegistryGetById(id string) (*types.Registry, error) {
	info, err := GetDataById[types.Registry](BucketRegistries, id)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeRegistryNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}

func RegistrysGetByNameAndUrl(name, url string, isLower bool) ([]*types.Registry, error) {
	registrys, err := GetDataAll[types.Registry](BucketRegistries)
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	var result []*types.Registry
	for _, registry := range registrys {
		if isLower && (strings.ToLower(registry.RegistryName) == strings.ToLower(name) || strings.ToLower(registry.URL) == strings.ToLower(url)) {
			result = append(result, registry)
		}
		if !isLower && (registry.RegistryName == name || registry.URL == url) {
			result = append(result, registry)
		}
	}
	return result, nil
}

func RegistryUpdate(List []*types.Registry) error {
	if err := TransactionUpdates(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketRegistries))
		if bucket == nil {
			return ErrBucketNotExist
		}
		for _, registry := range List {
			registryData, err := json.Marshal(registry)
			if err != nil {
				return fmt.Errorf("failed to encode user data: %s", err)
			}
			if err = bucket.Put([]byte(registry.RegistryId), registryData); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func RegistryDelete(id string) error {
	if err := DeleteData(BucketRegistries, id); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}
