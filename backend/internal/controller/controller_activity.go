package controller

import (
	"fmt"
	"log/slog"
	"slices"
	"time"

	"humpback/internal/db"
	"humpback/types"
)

var (
	ActivityCh = make(chan *ActivityEvent, 100)
)

type ActivityEvent struct {
	Data   *types.ActivityInfo
	Bucket string
}

func ReceiveActivities(stopCh <-chan struct{}) {
	defer close(ActivityCh)
	slog.Info("[Activity] Startup.")
	for {
		select {
		case <-stopCh:
			return
		case info := <-ActivityCh:
			if err := db.ActivityUpdate(info.Data, info.Bucket); err != nil {
				slog.Error("[Activity] Insert activity failed.", "Bucket", info.Bucket, "Id", info.Data.ActivityId, "Error", err)
			}
		}
	}
}

// InsertAccountActivity 构造account的activity，同时写入db
func InsertAccountActivity(oldInfo, currentInfo *types.User, action types.ActivityActoin, operateAt int64) {
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}
	description, oldContent, newContent := parseAccountActivityContent(oldInfo, currentInfo, action)
	data := &ActivityEvent{
		Bucket: db.ActivityBucketAcount,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s", operateAt, currentInfo.UserId, action),
			Action:       action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   currentInfo.UserId,
			Operator:     currentInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   currentInfo.UserId,
			ResourceName: currentInfo.Username,
		},
	}
	ActivityCh <- data
}

func parseAccountActivityContent(oldInfo, currentInfo *types.User, action types.ActivityActoin) (string, any, any) {
	switch action {
	case types.ActivityActionLogin:
		return "Log in to the system.", nil, nil
	case types.ActivityActionLogout:
		return "Log out of the system.", nil, nil
	case types.ActivityActionUpdate:
		var (
			oldContent map[string]any
			newContent map[string]any
		)
		if oldInfo != nil {
			oldContent = map[string]any{
				"username":    oldInfo.Username,
				"description": oldInfo.Description,
				"email":       oldInfo.Email,
				"phone":       oldInfo.Phone,
			}
		}
		if currentInfo != nil {
			newContent = map[string]any{
				"username":    currentInfo.Username,
				"description": currentInfo.Description,
				"email":       currentInfo.Email,
				"phone":       currentInfo.Phone,
			}
		}
		return "Update account information.", oldContent, newContent
	case types.ActivityActionChangePassword:
		return "Change the password.", nil, nil
	}
	return "", nil, nil
}

type ActivityUserInfo struct {
	OldUserInfo  *types.User
	NewUserInfo  *types.User
	OldTeams     []*types.Team
	NewTeams     []*types.Team
	Action       types.ActivityActoin
	OperatorInfo *types.User
	OperateAt    int64
}

// InsertUserActivity 构造user的activity，同时写入db
func InsertUserActivity(info *ActivityUserInfo) {
	var (
		operateAt                           = info.OperateAt
		description, oldContent, newContent = parseUserActivityContent(info)
		resourceInfo                        = info.NewUserInfo
	)
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}
	if info.Action == types.ActivityActionDelete {
		resourceInfo = info.OldUserInfo
	}

	data := &ActivityEvent{
		Bucket: db.ActivityBucketUsers,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action, resourceInfo.UserId),
			Action:       info.Action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   info.OperatorInfo.UserId,
			Operator:     info.OperatorInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   resourceInfo.UserId,
			ResourceName: resourceInfo.Username,
		},
	}
	ActivityCh <- data
}

func parseUserActivityContent(info *ActivityUserInfo) (string, any, any) {
	var (
		oldTeams    = make([]string, 0)
		newTeams    = make([]string, 0)
		oldUserRole = ""
		newUserRole = ""
		oldContent  map[string]any
		newContent  map[string]any
	)

	if info.OldUserInfo != nil {
		oldUserRole = types.UserRolesMap[info.OldUserInfo.Role]
		if len(info.OldTeams) > 0 {
			for _, teamId := range info.OldUserInfo.Teams {
				index := slices.IndexFunc(info.OldTeams, func(item *types.Team) bool {
					return item.TeamId == teamId
				})
				if index != -1 {
					oldTeams = append(oldTeams, info.OldTeams[index].Name)
				}
			}
		}
		oldContent = map[string]any{
			"username":    info.OldUserInfo.Username,
			"role":        oldUserRole,
			"description": info.OldUserInfo.Description,
			"email":       info.OldUserInfo.Email,
			"phone":       info.OldUserInfo.Phone,
			"teams":       oldTeams,
		}
	}
	if info.NewUserInfo != nil {
		newUserRole = types.UserRolesMap[info.NewUserInfo.Role]
		if len(info.NewTeams) > 0 {
			for _, teamId := range info.NewUserInfo.Teams {
				index := slices.IndexFunc(info.NewTeams, func(item *types.Team) bool {
					return item.TeamId == teamId
				})
				if index != -1 {
					newTeams = append(newTeams, info.NewTeams[index].Name)
				}
			}
		}
		newContent = map[string]any{
			"username":    info.NewUserInfo.Username,
			"role":        newUserRole,
			"description": info.NewUserInfo.Description,
			"email":       info.NewUserInfo.Email,
			"phone":       info.NewUserInfo.Phone,
			"teams":       newTeams,
		}
	}

	switch info.Action {
	case types.ActivityActionAdd:
		return fmt.Sprintf("Add user %s.", info.NewUserInfo.Username), nil, newContent
	case types.ActivityActionUpdate:
		return fmt.Sprintf("Update user %s.", info.NewUserInfo.Username), oldContent, newContent
	case types.ActivityActionDelete:
		return fmt.Sprintf("Delete user %s.", info.OldUserInfo.Username), oldContent, nil
	}
	return "", nil, nil
}

type ActivityTeamInfo struct {
	OldTeamInfo  *types.Team
	NewTeamInfo  *types.Team
	OldUsers     []*types.User
	NewUsers     []*types.User
	Action       types.ActivityActoin
	OperatorInfo *types.User
	OperateAt    int64
}

// InsertTeamActivity 构造team的activity，同时写入db
func InsertTeamActivity(info *ActivityTeamInfo) {
	var (
		operateAt                           = info.OperateAt
		description, oldContent, newContent = parseTeamActivityContent(info)
		resourceInfo                        = info.NewTeamInfo
	)
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}

	if info.Action == types.ActivityActionDelete {
		resourceInfo = info.OldTeamInfo
	}

	data := &ActivityEvent{
		Bucket: db.ActivityBucketTeams,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action, resourceInfo.TeamId),
			Action:       info.Action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   info.OperatorInfo.UserId,
			Operator:     info.OperatorInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   resourceInfo.TeamId,
			ResourceName: resourceInfo.Name,
		},
	}
	ActivityCh <- data
}

func parseTeamActivityContent(info *ActivityTeamInfo) (string, any, any) {
	var (
		oldUsers   = make([]string, 0)
		newUsers   = make([]string, 0)
		oldContent map[string]any
		newContent map[string]any
	)

	if info.OldTeamInfo != nil {
		if len(info.OldUsers) > 0 {
			for _, userId := range info.OldTeamInfo.Users {
				index := slices.IndexFunc(info.OldUsers, func(item *types.User) bool {
					return item.UserId == userId
				})
				if index != -1 {
					oldUsers = append(oldUsers, info.OldUsers[index].Username)
				}
			}
		}
		oldContent = map[string]any{
			"name":        info.OldTeamInfo.Name,
			"description": info.OldTeamInfo.Description,
			"users":       oldUsers,
		}
	}
	if info.NewTeamInfo != nil {
		if len(info.NewUsers) > 0 {
			for _, userId := range info.NewTeamInfo.Users {
				index := slices.IndexFunc(info.NewUsers, func(item *types.User) bool {
					return item.UserId == userId
				})
				if index != -1 {
					newUsers = append(newUsers, info.NewUsers[index].Username)
				}
			}
		}
		newContent = map[string]any{
			"name":        info.NewTeamInfo.Name,
			"description": info.NewTeamInfo.Description,
			"users":       newUsers,
		}
	}

	switch info.Action {
	case types.ActivityActionAdd:
		return fmt.Sprintf("Add team %s.", info.NewTeamInfo.Name), nil, newContent
	case types.ActivityActionUpdate:
		return fmt.Sprintf("Update team %s.", info.NewTeamInfo.Name), oldContent, newContent
	case types.ActivityActionDelete:
		return fmt.Sprintf("Delete team %s.", info.OldTeamInfo.Name), oldContent, nil
	}
	return "", nil, nil
}

type ActivityNodeInfo struct {
	OldNodeInfo  *types.Node
	NewNodeInfo  *types.Node
	Action       types.ActivityActoin
	OperatorInfo *types.User
	OperateAt    int64
}

// InsertNodeActivity 构造node的activity，同时写入db
func InsertNodeActivity(info *ActivityNodeInfo) {
	var (
		operateAt                           = info.OperateAt
		description, oldContent, newContent = parseNodeActivityContent(info)
		resourceInfo                        = info.NewNodeInfo
	)
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}

	if info.Action == types.ActivityActionDelete {
		resourceInfo = info.OldNodeInfo
	}

	data := &ActivityEvent{
		Bucket: db.ActivityBucketNodes,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action, resourceInfo.NodeId),
			Action:       info.Action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   info.OperatorInfo.UserId,
			Operator:     info.OperatorInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   resourceInfo.NodeId,
			ResourceName: resourceInfo.IpAddress,
		},
	}
	ActivityCh <- data
}

func parseNodeActivityContent(info *ActivityNodeInfo) (string, any, any) {
	var (
		oldContent map[string]any
		newContent map[string]any
	)

	if info.OldNodeInfo != nil {
		oldContent = map[string]any{
			"ipAddress": info.OldNodeInfo.IpAddress,
			"labels":    info.OldNodeInfo.Labels,
		}
	}
	if info.NewNodeInfo != nil {
		newContent = map[string]any{
			"ipAddress": info.NewNodeInfo.IpAddress,
			"labels":    info.NewNodeInfo.Labels,
		}
	}

	switch info.Action {
	case types.ActivityActionAdd:
		return fmt.Sprintf("Add node %s.", info.NewNodeInfo.IpAddress), nil, nil
	case types.ActivityActionUpdateLabel:
		return fmt.Sprintf("Update node %s.", info.NewNodeInfo.IpAddress), oldContent, newContent
	case types.ActivityActionDelete:
		return fmt.Sprintf("Delete node %s.", info.OldNodeInfo.IpAddress), oldContent, nil
	case types.ActivityActionEnable:
		return fmt.Sprintf("Enable node %s.", info.NewNodeInfo.IpAddress), nil, nil
	case types.ActivityActionDisable:
		return fmt.Sprintf("Disable node %s.", info.NewNodeInfo.IpAddress), nil, nil
	}
	return "", nil, nil
}

type ActivityRegistryInfo struct {
	OldRegistryInfo *types.Registry
	NewRegistryInfo *types.Registry
	Action          types.ActivityActoin
	OperatorInfo    *types.User
	OperateAt       int64
}

// InsertRegistryActivity 构造registry的activity，同时写入db
func InsertRegistryActivity(info *ActivityRegistryInfo) {
	var (
		operateAt                           = info.OperateAt
		description, oldContent, newContent = parseRegistryActivityContent(info)
		resourceInfo                        = info.NewRegistryInfo
	)
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}

	if info.Action == types.ActivityActionDelete {
		resourceInfo = info.OldRegistryInfo
	}

	data := &ActivityEvent{
		Bucket: db.ActivityBucketRegistries,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action, resourceInfo.RegistryId),
			Action:       info.Action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   info.OperatorInfo.UserId,
			Operator:     info.OperatorInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   resourceInfo.RegistryId,
			ResourceName: resourceInfo.URL,
		},
	}
	ActivityCh <- data
}

func parseRegistryActivityContent(info *ActivityRegistryInfo) (string, any, any) {
	var (
		oldContent map[string]any
		newContent map[string]any
	)

	if info.OldRegistryInfo != nil {
		oldContent = map[string]any{
			"url":       info.OldRegistryInfo.URL,
			"username":  info.OldRegistryInfo.Username,
			"isDefault": info.OldRegistryInfo.IsDefault,
		}
	}
	if info.NewRegistryInfo != nil {
		newContent = map[string]any{
			"url":       info.NewRegistryInfo.URL,
			"username":  info.NewRegistryInfo.Username,
			"isDefault": info.NewRegistryInfo.IsDefault,
		}
	}

	switch info.Action {
	case types.ActivityActionAdd:
		return fmt.Sprintf("Add registry %s.", info.NewRegistryInfo.URL), nil, newContent
	case types.ActivityActionUpdate:
		return fmt.Sprintf("Update registry %s.", info.NewRegistryInfo.URL), oldContent, newContent
	case types.ActivityActionDelete:
		return fmt.Sprintf("Delete registry %s.", info.OldRegistryInfo.URL), oldContent, nil
	}
	return "", nil, nil
}

type ActivityConfigInfo struct {
	OldConfigInfo *types.Config
	NewConfigInfo *types.Config
	Action        types.ActivityActoin
	OperatorInfo  *types.User
	OperateAt     int64
}

// InsertConfigActivity 构造registry的activity，同时写入db
func InsertConfigActivity(info *ActivityConfigInfo) {
	var (
		operateAt                           = info.OperateAt
		description, oldContent, newContent = parseConfigActivityContent(info)
		resourceInfo                        = info.NewConfigInfo
	)
	if operateAt == 0 {
		operateAt = time.Now().UnixMilli()
	}

	if info.Action == types.ActivityActionDelete {
		resourceInfo = info.OldConfigInfo
	}

	data := &ActivityEvent{
		Bucket: db.ActivityBucketConfigs,
		Data: &types.ActivityInfo{
			ActivityId:   fmt.Sprintf("%d-%s-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action, resourceInfo.ConfigId),
			Action:       info.Action,
			Description:  description,
			OldContent:   oldContent,
			NewContent:   newContent,
			OperatorId:   info.OperatorInfo.UserId,
			Operator:     info.OperatorInfo.Username,
			OperateAt:    operateAt,
			ResourceId:   resourceInfo.ConfigId,
			ResourceName: resourceInfo.ConfigName,
		},
	}
	ActivityCh <- data
}

func parseConfigActivityContent(info *ActivityConfigInfo) (string, any, any) {
	var (
		oldContent map[string]any
		newContent map[string]any
	)

	if info.OldConfigInfo != nil {
		oldContent = map[string]any{
			"configName":  info.OldConfigInfo.ConfigName,
			"description": info.OldConfigInfo.Description,
			"configType":  types.ConfigTypesMap[info.OldConfigInfo.ConfigType],
			"configValue": info.OldConfigInfo.ConfigValue,
		}
	}
	if info.NewConfigInfo != nil {
		newContent = map[string]any{
			"configName":  info.NewConfigInfo.ConfigName,
			"description": info.NewConfigInfo.Description,
			"configType":  types.ConfigTypesMap[info.NewConfigInfo.ConfigType],
			"configValue": info.NewConfigInfo.ConfigValue,
		}
	}

	switch info.Action {
	case types.ActivityActionAdd:
		return fmt.Sprintf("Add config %s.", info.NewConfigInfo.ConfigName), nil, newContent
	case types.ActivityActionUpdate:
		return fmt.Sprintf("Update config %s.", info.NewConfigInfo.ConfigName), oldContent, newContent
	case types.ActivityActionDelete:
		return fmt.Sprintf("Delete config %s.", info.OldConfigInfo.ConfigName), oldContent, nil
	}
	return "", nil, nil
}
