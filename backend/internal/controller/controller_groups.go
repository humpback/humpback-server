package controller

import (
	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/types"
)

func GroupCreate(reqInfo *models.GroupCreateReqInfo) (string, error) {
	if err := groupCheckNameExist(reqInfo.GroupName, ""); err != nil {
		return "", err
	}
	newInfo := reqInfo.NewGroupInfo()
	if err := db.GroupUpdate(newInfo); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return newInfo.GroupId, nil
}

func GroupUpdate(userInfo *types.User, reqInfo *models.GroupUpdateReqInfo) (string, error) {
	oldInfo, err := Group(userInfo, reqInfo.GroupId)
	if err != nil {
		return "", err
	}
	if !userInfo.InGroup(oldInfo) {
		return "", response.NewBadRequestErr(locales.CodeGroupNoPermission)
	}
	if err := groupCheckNameExist(reqInfo.GroupName, reqInfo.GroupId); err != nil {
		return "", err
	}
	newInfo := reqInfo.NewGroupInfo(oldInfo)
	if err = db.GroupUpdate(newInfo); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return newInfo.GroupId, err
}

func groupCheckNameExist(name, id string) error {
	sameNames, err := db.GroupsGetByName(name, true)
	if err != nil {
		return response.NewRespServerErr(err.Error())
	}
	for _, info := range sameNames {
		if info.GroupId != id {
			return response.NewBadRequestErr(locales.CodeGroupNameAlreadyExist)
		}
	}
	return nil
}

func GroupUpdateNodes(userInfo *types.User, info *models.GroupUpdateNodesReqInfo) (string, error) {
	oldInfo, err := Group(userInfo, info.GroupId)
	if err != nil {
		return "", err
	}
	if !userInfo.InGroup(oldInfo) {
		return "", response.NewBadRequestErr(locales.CodeGroupNoPermission)
	}

	if !info.IsDelete {
		_, err := db.NodesGetByIds(info.Nodes, false)
		if err != nil {
			if err == db.ErrKeyNotExist {
				return "", response.NewBadRequestErr(locales.CodeNodesNotExist)
			}
			return "", response.NewRespServerErr(err.Error())
		}
	}
	newGroup := info.NewGroupInfo(oldInfo)
	if err = db.GroupUpdate(newGroup); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	//todo 往scheduler发送消息
	return info.GroupId, nil
}

func Group(userInfo *types.User, id string) (*types.NodesGroups, error) {
	info, err := db.GroupGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	if !userInfo.InGroup(info) {
		return nil, response.NewBadRequestErr(locales.CodeGroupNoPermission)
	}
	return info, nil
}

func GroupQuery(queryInfo *models.GroupQueryReqInfo) (*response.QueryResult[types.NodesGroups], error) {
	groups, err := db.GroupsGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := queryInfo.QueryFilter(groups)
	return response.NewQueryResult[types.NodesGroups](
		len(result),
		types.QueryPagination[types.NodesGroups](queryInfo.PageInfo, result),
	), nil
}

func GroupDelete(id string) error {
	_, err := db.GroupGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			//todo 往scheduler发送消息
			return nil
		}
		return response.NewRespServerErr(err.Error())
	}
	if err = db.GroupDelete(id); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	//todo 往scheduler发送消息
	return nil
}

func GroupNodesQuery(groupId string, userInfo *types.User, queryInfo *models.GroupQueryNodesReqInfo) (*response.QueryResult[types.Node], error) {
	groupInfo, err := Group(userInfo, groupId)
	if err != nil {
		return nil, err
	}
	if !userInfo.InGroup(groupInfo) {
		return nil, response.NewBadRequestErr(locales.CodeGroupNoPermission)
	}
	nodes, err := NodesGetByIds(groupInfo.Nodes, true)
	if err != nil {
		return nil, err
	}
	result := queryInfo.QueryFilter(nodes)
	return response.NewQueryResult[types.Node](
		len(result),
		types.QueryPagination[types.Node](queryInfo.PageInfo, result),
	), nil
}
