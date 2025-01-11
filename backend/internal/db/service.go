package db

import "humpback/types"

func GetAllService() ([]*types.Service, error) {
	return GetDataAll[types.Service](BucketServices)
}

func GetServiceById(serviceId string) (*types.Service, error) {
	return GetDataById[types.Service](BucketServices, serviceId)
}

func SaveService(data *types.Service) error {
	return SaveData(BucketServices, data.ServiceId, data)
}
