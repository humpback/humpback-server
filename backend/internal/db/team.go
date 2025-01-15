package db

import (
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func TeamsByIds(ids []string, ingonreNotExist bool) ([]*types.Team, error) {
	teams, err := GetDataByIds[types.Team](BucketTeams, ids, ingonreNotExist)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeTeamNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return teams, nil
}

func TeamById(id string) (*types.Team, error) {
	info, err := GetDataById[types.Team](BucketTeams, id)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeTeamNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}
