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
    slog.Info("[Activiy] Startup.")
    for {
        select {
        case <-stopCh:
            return
        case info := <-ActivityCh:
            if err := db.ActivityUpdate(info.Data, info.Bucket); err != nil {
                slog.Error("[Activiy] Insert acitivity failed.", "Bucket", info.Bucket, "Id", info.Data.ActivityId, "Error", err)
            }
        }
    }
}

func InsertAccountAcitvity(oldInfo, currentInfo *types.User, action types.ActivityActoin, operateAt int64) {
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
    case types.ActivityActoinLogin:
        return "Log in to the system.", nil, nil
    case types.ActivityActoinLogout:
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

func InsertUserActivity(info *ActivityUserInfo) {
    operateAt := info.OperateAt
    if operateAt == 0 {
        operateAt = time.Now().UnixMilli()
    }
    description, oldContent, newContent := parseUserActivityContent(info)
    data := &ActivityEvent{
        Bucket: db.ActivityBucketAcount,
        Data: &types.ActivityInfo{
            ActivityId:   fmt.Sprintf("%d-%s-%s", operateAt, info.OperatorInfo.UserId, info.Action),
            Action:       info.Action,
            Description:  description,
            OldContent:   oldContent,
            NewContent:   newContent,
            OperatorId:   info.OperatorInfo.UserId,
            Operator:     info.OperatorInfo.Username,
            OperateAt:    operateAt,
            ResourceId:   info.NewUserInfo.UserId,
            ResourceName: info.NewUserInfo.Username,
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
    case types.ActivityActionCreate:
        return fmt.Sprintf("Add user %s", info.NewUserInfo.Username), nil, newContent
    case types.ActivityActionUpdate:
        return fmt.Sprintf("Update user %s", info.NewUserInfo.Username), oldContent, newContent
    case types.ActivityActionDelete:
        return fmt.Sprintf("Delete user %s", info.OldUserInfo.Username), oldContent, nil
    }
    return "", nil, nil
}
