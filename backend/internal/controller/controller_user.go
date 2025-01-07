package controller

import (
	"humpback/api/handle/models"
	"humpback/types"
)

func UserLogin(info *models.UserLoginReqInfo) (*types.User, error) {
	//userInfo, err := db.UserGetByNamePsd(info.Name, info.Password)
	//if err != nil {
	//	return nil, response.NewRespServerErr(err.Error())
	//}
	//if userInfo == nil {
	//	return nil, response.NewBadRequestErr()
	//}
	return nil, nil
}
