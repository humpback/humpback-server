package db

import "humpback/types"

func ServicesGetAll() ([]*types.Service, error) {
	return GetDataAll[types.Service](BucketServices)
}

func ServicesGetByPrefix(prefix string) ([]*types.Service, error) {
	return GetDataByPrefix[types.Service](BucketServices, prefix)
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
