package controller

import (
	"fmt"
	"log/slog"
	"slices"

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
	slog.Info("[Supper Admin] Account check start...")
	adminConfig := config.AdminArgs()
	user, err := db.UserGetSupperAdmin()
	if err != nil {
		return fmt.Errorf("Check admin account failed: %s", err)
	}
	if user == nil {
		var (
			t  = utils.NewActionTimestamp()
			id = utils.NewGuidStr()
		)
		if err = db.UserUpdate(id, &types.User{
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
	slog.Info("[Supper Admin] Account check completed.")
	return nil
}

func UserLogin(reqInfo *models.UserLoginReqInfo) (*types.User, string, error) {
	user, err := userCheckNamePsd(reqInfo.Username, reqInfo.Password)
	if err != nil {
		return nil, "", err
	}
	sessionInfo := &types.Session{
		SessionId: utils.NewGuidStr(),
		UserId:    user.UserId,
	}
	if err = SessionUpdate(sessionInfo); err != nil {
		return nil, "", err
	}
	return user, sessionInfo.SessionId, nil
}

func userCheckNamePsd(username, password string) (*types.User, error) {
	users, err := Users()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			if user.Password == password {
				return user, nil
			}
			return nil, response.NewBadRequestErr(locales.CodePasswordIsWrong)
		}
	}
	return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
}

func MeUpdate(userInfo *types.User) error {
	if err := db.UserUpdate(userInfo.UserId, userInfo); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func MeChangePassword(userInfo *types.User, reqInfo *models.MeChangePasswordReqInfo) error {
	if userInfo.Password != reqInfo.OldPassword {
		return response.NewBadRequestErr(locales.CodeOldPasswordIsWrong)
	}
	userInfo.Password = reqInfo.NewPassword
	userInfo.UpdatedAt = utils.NewActionTimestamp()
	if err := MeUpdate(userInfo); err != nil {
		return err
	}
	if err := SessionDeleteByUserId(userInfo.UserId); err != nil {
		return err
	}
	return nil
}

func UserCreate(reqInfo *models.UserCreateReqInfo) (string, error) {
	err := userCheckNameExist(reqInfo.Username, "")
	if err != nil {
		return "", err
	}
	var (
		teams    = make([]*types.Team, 0)
		userInfo = reqInfo.NewUserInfo()
	)
	if len(reqInfo.Teams) > 0 {
		teams, err = TeamsGetByIds(reqInfo.Teams, false)
		if err != nil {
			return "", err
		}
	}
	for _, team := range teams {
		team.Users = append(team.Users, userInfo.UserId)
	}
	id, err := db.UserUpdateAndTeams(userInfo, teams)
	if err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return id, nil
}

func UserUpdate(reqInfo *models.UserUpdateReqInfo, operator *types.User) (string, error) {
	oldUser, err := userCheckRole(reqInfo, operator)
	if err != nil {
		return "", err
	}
	if err = userCheckNameExist(reqInfo.Username, reqInfo.UserId); err != nil {
		return "", err
	}
	newUser, clearSession := reqInfo.NewUserInfo(oldUser)
	updateTeams, err := userUpdateCheckTeams(oldUser.Teams, newUser.Teams, newUser.UserId)
	if err != nil {
		return "", err
	}

	id, err := db.UserUpdateAndTeams(newUser, updateTeams)
	if err != nil {
		return "", response.NewRespServerErr(err.Error())
	}

	if clearSession {
		if err = SessionDeleteByUserId(newUser.UserId); err != nil {
			return "", err
		}
	}
	return id, nil
}

func userCheckRole(reqInfo *models.UserUpdateReqInfo, operator *types.User) (*types.User, error) {
	user, err := User(reqInfo.UserId)
	if err != nil {
		return nil, err
	}
	if types.IsAdmin(user.Role) && !types.IsSupperAdmin(operator.Role) {
		return nil, response.NewRespServerErr(locales.CodeNoPermission)
	}
	return user, nil
}

func userCheckNameExist(username, userId string) error {
	users, err := db.UsersGetByName(username, true)
	if err != nil {
		return response.NewRespServerErr(err.Error())
	}
	for _, user := range users {
		if user.UserId != userId {
			return response.NewBadRequestErr(locales.CodeUserNameAlreadyExist)
		}
	}
	return nil
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
	teamList, err := TeamsGetByIds(maps.Keys(teamIdMap), false)
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

func User(id string) (*types.User, error) {
	info, err := db.UserGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}

func Users() ([]*types.User, error) {
	users, err := db.UsersGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	return users, nil
}

func UsersQuery(queryInfo *models.UserQueryReqInfo) (*response.QueryResult[types.User], error) {
	users, err := Users()
	if err != nil {
		return nil, err
	}
	result := queryInfo.QueryFilter(users)
	return response.NewQueryResult[types.User](
		len(result),
		types.QueryPagination[types.User](queryInfo.PageInfo, result),
	), nil
}

func UsersGetByIds(ids []string, ignoreNotExist bool) ([]*types.User, error) {
	users, err := db.UsersGetByIds(ids, ignoreNotExist)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeUserNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return users, nil
}

func UsersGetByTeamId(teamId string) ([]*types.User, error) {
	teamInfo, err := Team(teamId)
	if err != nil {
		return nil, err
	}
	users, err := UsersGetByIds(teamInfo.Users, true)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		user.Password = ""
	}
	return users, nil
}

func UserDelete(id string, operator *types.User) error {
	info, err := db.UserGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil
		}
		return response.NewRespServerErr(err.Error())
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
		return response.NewRespServerErr(err.Error())
	}
	if err = db.SessionBatchDeleteByUserId(id); err != nil {
		slog.Warn("[User Delete] Clear session failed.", "UserId", id, "UserName", info.Username, "Error", err.Error())
	}
	return nil
}

func userDeleteCheckTeams(teams []string, userId string) ([]*types.Team, error) {
	if len(teams) == 0 {
		return nil, nil
	}
	teamList, err := TeamsGetByIds(teams, true)
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
