package controller

import (
	"fmt"
	"log/slog"
	"slices"
	"time"

	"golang.org/x/exp/maps"
	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/config"
	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"
)

func InitAdminUser() error {
	adminConfig := config.AdminArgs()
	user, err := db.UserFindSupperAdmin()
	if err != nil {
		return fmt.Errorf("Check admin account failed: %s", err)
	}
	if user == nil {
		var (
			t  = time.Now().UnixMilli()
			id = utils.NewGuidStr()
		)
		if err = db.UserInit(id, &types.User{
			UserId:    id,
			Username:  adminConfig.Username,
			Email:     "",
			Password:  adminConfig.Password,
			Phone:     "",
			Role:      types.UserRoleSupperAdmin,
			CreatedAt: t,
			UpdatedAt: t,
			Teams:     nil,
		}); err != nil {
			return fmt.Errorf("Create admin account failed: %s", err)
		}
	}
	slog.Info("Admin account check success")
	return nil
}

func UserLogin(info *models.UserLoginReqInfo) (*types.User, string, error) {
	userInfo, err := db.UserGetByNamePsd(info.Username, info.Password)
	if err != nil {
		return nil, "", err
	}
	sessionInfo := &types.Session{
		SessionId: utils.NewGuidStr(),
		UserId:    userInfo.UserId,
	}
	if err = SessionUpdate(sessionInfo); err != nil {
		return nil, "", err
	}
	return userInfo, sessionInfo.SessionId, nil
}

func MeUpdate(userInfo *types.User) error {
	if err := db.MeUpdate(userInfo.UserId, userInfo); err != nil {
		return err
	}
	return nil
}

func MeChangePassword(userInfo *types.User, reqInfo *models.MeChangePasswordReqInfo) error {
	if userInfo.Password != reqInfo.OldPassword {
		return response.NewBadRequestErr(locales.CodeOldPasswordIsWrong)
	}
	userInfo.Password = reqInfo.NewPassword
	userInfo.UpdatedAt = time.Now().UnixMilli()
	if err := db.MeUpdate(userInfo.UserId, userInfo); err != nil {
		return err
	}
	if err := db.SessionBatchDeleteByUserId(userInfo.UserId); err != nil {
		return err
	}
	return nil
}

func UserCreate(info *models.UserCreateReqInfo) (string, error) {
	oldInfo, err := db.UserGetByName(info.Username)
	if err != nil {
		return "", err
	}
	if oldInfo != nil {
		return "", response.NewBadRequestErr(locales.CodeUserAlreadyExist)
	}
	var (
		teams    = make([]*types.Team, 0)
		userInfo = info.NewUserInfo()
	)
	if len(info.Teams) > 0 {
		teams, err = db.TeamsByIds(info.Teams, false)
		if err != nil {
			return "", err
		}
	}
	for _, team := range teams {
		team.Users = append(team.Users, userInfo.UserId)
	}
	id, err := db.UserUpdateAndTeams(userInfo, teams)
	if err != nil {
		return "", err
	}
	return id, nil
}

func UserUpdate(info *models.UserUpdateReqInfo, operator *types.User) (string, error) {
	oldUserInfo, err := db.UserGetById(info.UserId)
	if err != nil {
		return "", err
	}
	if oldUserInfo.Role == types.UserRoleAdmin && operator.Role != types.UserRoleSupperAdmin {
		return "", response.NewRespServerErr(locales.CodeNoPermission)
	}
	newUserInfo, clearSession := info.NewUserInfo(oldUserInfo)
	updateTeams, err := userUpdateCheckTeams(oldUserInfo.Teams, newUserInfo.Teams, newUserInfo.UserId)
	if err != nil {
		return "", err
	}

	id, err := db.UserUpdateAndTeams(newUserInfo, updateTeams)
	if err != nil {
		return "", err
	}

	if clearSession {
		if err = db.SessionBatchDeleteByUserId(newUserInfo.UserId); err != nil {
			return "", err
		}
	}
	return id, nil
}

func userUpdateCheckTeams(oldTeams, newTeams []string, userId string) ([]*types.Team, error) {
	if len(oldTeams) == 0 && len(newTeams) == 0 {
		return nil, nil
	}

	var teamIdMap = map[string]int{}
	for _, teamId := range oldTeams {
		teamIdMap[teamId] = -1
	}
	for _, teamId := range newTeams {
		if _, ok := teamIdMap[teamId]; ok {
			teamIdMap[teamId] = 0
		} else {
			teamIdMap[teamId] = 1
		}
	}
	teamList, err := db.TeamsByIds(maps.Keys(teamIdMap), false)
	if err != nil {
		return nil, err
	}
	var resultTeams = make([]*types.Team, 0)
	for _, team := range teamList {
		action, ok := teamIdMap[team.TeamId]
		if !ok && action == 0 {
			continue
		}
		index := slices.Index(team.Users, userId)
		if action == -1 && index != -1 {
			team.Users = append(team.Users[:index], team.Users[index+1:]...)
			resultTeams = append(resultTeams, team)
			continue
		}
		if action == 1 && index == -1 {
			team.Users = append(team.Users, userId)
			resultTeams = append(resultTeams, team)
		}
	}

	return resultTeams, nil
}

func UserQuery(queryInfo *models.UserQueryReqInfo) (*types.QueryResult[types.User], error) {
	users, err := db.UserGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := make([]*types.User, 0)
	for _, user := range users {
		user.Password = ""
		if queryInfo.QueryFilter(user) {
			result = append(result, user)
		}
	}
	queryInfo.QuerySort(result)
	return types.NewQueryResult[types.User](
		len(result),
		types.QueryPagination[types.User](queryInfo.PageInfo, result),
	), nil

}

func User(id string) (*types.User, error) {
	info, err := db.UserGetById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func UserDelete(id string, operator *types.User) error {
	info, err := db.UserGetById(id)
	if err != nil {
		e := err.(*response.ErrInfo)
		if e.Code == locales.CodeUserNotExist {
			return nil
		}
		return err
	}
	if id == operator.UserId {
		return response.NewBadRequestErr(locales.CodeUserIsOwner)
	}
	if info.Role == types.UserRoleSupperAdmin || (info.Role == types.UserRoleAdmin && operator.Role != types.UserRoleSupperAdmin) {
		return response.NewBadRequestErr(locales.CodeNoPermission)
	}

	teams, err := userDeleteCheckTeams(info.Teams, id)
	if err != nil {
		return err
	}
	if err = db.UserDelete(id, teams); err != nil {
		return err
	}
	if err = db.SessionBatchDeleteByUserId(id); err != nil {
		slog.Warn("[User Delete] clear session failed.", "UserId", id, "UserName", info.Username, "Error", err.Error())
	}
	return nil
}

func userDeleteCheckTeams(teams []string, userId string) ([]*types.Team, error) {
	if len(teams) == 0 {
		return nil, nil
	}
	teamList, err := db.TeamsByIds(teams, true)
	if err != nil {
		return nil, err
	}
	var result []*types.Team
	for _, team := range teamList {
		if index := slices.Index(team.Users, userId); index != -1 {
			team.Users = append(team.Users[:index], team.Users[index+1:]...)
			result = append(result, team)
		}
	}
	return result, nil
}

func UsersByTeamId(teamId string) ([]*types.User, error) {
	teamInfo, err := db.TeamById(teamId)
	if err != nil {
		return nil, err
	}
	users, err := db.UsersByIds(teamInfo.Users, true)
	for _, user := range users {
		user.Password = ""
	}
	return users, nil
}
