package db

import (
	"humpback/types"
)

func ServicesGetAll() ([]*types.Service, error) {
	return GetDataAll[types.Service](BucketServices)
}

func ServicesGetValidByPrefix(prefix string) ([]*types.Service, error) {
	services, err := GetDataByPrefix[types.Service](BucketServices, prefix)
	if err != nil {
		return nil, err
	}
	result := make([]*types.Service, 0)
	for _, service := range services {
		if !service.IsDelete {
			result = append(result, service)
		}
	}
	return result, nil
}

func ServiceGetTotalByPrefix(prefix string) (int, error) {
	return GetDataTotalByPrefix[types.Service](BucketServices, prefix)
}

func ServiceGetById(serviceId string) (*types.Service, error) {
	return GetDataById[types.Service](BucketServices, serviceId)
}

func ServiceUpdate(data *types.Service) error {
	return SaveData(BucketServices, data.ServiceId, data)
}

func ServiceDelete(id string) error {
	return DeleteData(BucketServices, id)
}
