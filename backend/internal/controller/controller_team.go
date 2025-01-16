package controller

import (
	"slices"

	"golang.org/x/exp/maps"

	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/types"
)

func TeamCreate(info *models.TeamCreateReqInfo) (string, error) {
	err := teamCreateCheckName(info)
	if err != nil {
		return "", err
	}
	var (
		users    = make([]*types.User, 0)
		teamInfo = info.NewTeamInfo()
	)
	if len(info.Users) > 0 {
		users, err = db.UsersQueryByIds(info.Users, false)
		if err != nil {
			return "", err
		}
	}
	for _, user := range users {
		user.Teams = append(user.Teams, teamInfo.TeamId)
	}
	id, err := db.TeamUpdateAndUsers(teamInfo, users)
	if err != nil {
		return "", err
	}
	return id, nil
}

func teamCreateCheckName(info *models.TeamCreateReqInfo) error {
	sameNameTeams, err := db.TeamsGetByName(info.Name)
	if err != nil {
		return err
	}
	if len(sameNameTeams) > 0 {
		return response.NewBadRequestErr(locales.CodeTeamAlreadyExist)
	}
	return nil
}

func TeamUpdate(info *models.TeamUpdateReqInfo) (string, error) {
	if err := teamUpdateCheckName(info); err != nil {
		return "", err
	}

	oldTeam, err := db.TeamGetById(info.TeamId)
	if err != nil {
		return "", err
	}

	newTeamInfo := info.NewTeamInfo(oldTeam)
	updateUsers, err := teamUpdateCheckUsers(oldTeam.Users, newTeamInfo.Users, newTeamInfo.TeamId)
	if err != nil {
		return "", err
	}

	id, err := db.TeamUpdateAndUsers(newTeamInfo, updateUsers)
	if err != nil {
		return "", err
	}
	return id, nil
}

func teamUpdateCheckName(info *models.TeamUpdateReqInfo) error {
	sameNameTeams, err := db.TeamsGetByName(info.Name)
	if err != nil {
		return err
	}

	if len(sameNameTeams) > 1 || len(sameNameTeams) == 1 && sameNameTeams[0].TeamId != info.TeamId {
		return response.NewBadRequestErr(locales.CodeTeamAlreadyExist)
	}
	return nil
}

func teamUpdateCheckUsers(oldUsers, newUsers []string, teamId string) ([]*types.User, error) {
	if len(oldUsers) == 0 && len(newUsers) == 0 {
		return nil, nil
	}

	var userIdMap = map[string]int{}
	for _, userId := range oldUsers {
		userIdMap[userId] = -1
	}
	for _, userId := range newUsers {
		if _, ok := userIdMap[userId]; ok {
			userIdMap[userId] = 0
		} else {
			userIdMap[userId] = 1
		}
	}
	userList, err := db.UsersQueryByIds(maps.Keys(userIdMap), false)
	if err != nil {
		return nil, err
	}
	var resultUsers = make([]*types.User, 0)
	for _, user := range userList {
		action, ok := userIdMap[user.UserId]
		if !ok && action == 0 {
			continue
		}
		index := slices.Index(user.Teams, teamId)
		if action == -1 && index != -1 {
			user.Teams = append(user.Teams[:index], user.Teams[index+1:]...)
			resultUsers = append(resultUsers, user)
			continue
		}
		if action == 1 && index == -1 {
			user.Teams = append(user.Teams, teamId)
			resultUsers = append(resultUsers, user)
		}
	}
	return resultUsers, nil
}

func Team(id string) (*types.Team, error) {
	info, err := db.TeamGetById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func TeamsByUserId(userId string) ([]*types.Team, error) {
	userInfo, err := db.UserGetById(userId)
	if err != nil {
		return nil, err
	}
	teams, err := db.TeamsQueryByIds(userInfo.Teams, true)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func TeamQuery(queryInfo *models.TeamQueryReqInfo) (*response.QueryResult[types.Team], error) {
	users, err := db.TeamGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := queryInfo.QueryFilter(users)
	return response.NewQueryResult[types.Team](
		len(result),
		types.QueryPagination[types.Team](queryInfo.PageInfo, result),
	), nil
}

func TeamDelete(id string) error {
	info, err := db.TeamGetById(id)
	if err != nil {
		e := err.(*response.ErrInfo)
		if e.Code == locales.CodeTeamNotExist {
			return nil
		}
		return err
	}

	users, err := teamDeleteCheckUsers(info.Users, id)
	if err != nil {
		return err
	}
	if err = db.TeamDelete(id, users); err != nil {
		return err
	}
	return nil
}

func teamDeleteCheckUsers(users []string, teamId string) ([]*types.User, error) {
	if len(users) == 0 {
		return nil, nil
	}
	userList, err := db.UsersQueryByIds(users, true)
	if err != nil {
		return nil, err
	}
	var result []*types.User
	for _, user := range userList {
		if index := slices.Index(user.Teams, teamId); index != -1 {
			user.Teams = append(user.Teams[:index], user.Teams[index+1:]...)
			result = append(result, user)
		}
	}
	return result, nil
}
