package controller

import (
	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/types"
)

func ConfigCreate(reqInfo *models.ConfigCreateReqInfo) (string, error) {
	err := configCreateCheckName(reqInfo)
	if err != nil {
		return "", err
	}
	newInfo := reqInfo.NewConfigInfo()
	if err = db.ConfigUpdate(newInfo); err != nil {
		return "", err
	}
	return newInfo.ConfigId, err
}

func configCreateCheckName(reqInfo *models.ConfigCreateReqInfo) error {
	sameNames, err := db.ConfigsGetByName(reqInfo.ConfigName, true)
	if err != nil {
		return err
	}
	if len(sameNames) > 0 {
		return response.NewBadRequestErr(locales.CodeConfigNameAlreadyExist)
	}
	return nil
}

func ConfigUpdate(reqInfo *models.ConfigUpdateReqInfo) (string, error) {
	if err := configUpdateCheckName(reqInfo); err != nil {
		return "", err
	}
	oldInfo, err := db.ConfigGetById(reqInfo.ConfigId)
	if err != nil {
		return "", err
	}
	newInfo := reqInfo.NewConfigInfo(oldInfo)
	if err = db.ConfigUpdate(newInfo); err != nil {
		return "", err
	}
	return newInfo.ConfigId, err
}

func configUpdateCheckName(reqInfo *models.ConfigUpdateReqInfo) error {
	sameNames, err := db.ConfigsGetByName(reqInfo.ConfigName, true)
	if err != nil {
		return err
	}
	if len(sameNames) > 1 || len(sameNames) == 1 && sameNames[0].ConfigId != reqInfo.ConfigId {
		return response.NewBadRequestErr(locales.CodeConfigNameAlreadyExist)
	}
	return nil
}

func Config(id string) (*types.Config, error) {
	info, err := db.ConfigGetById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func ConfigQuery(queryInfo *models.ConfigQueryReqInfo) (*response.QueryResult[types.Config], error) {
	configs, err := db.ConfigGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := queryInfo.QueryFilter(configs)
	return response.NewQueryResult[types.Config](
		len(result),
		types.QueryPagination[types.Config](queryInfo.PageInfo, result),
	), nil
}

func ConfigDelete(id string) error {
	if err := db.ConfigDelete(id); err != nil {
		return err
	}
	return nil
}
