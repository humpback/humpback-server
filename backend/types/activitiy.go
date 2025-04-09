package types

type ActivityActoin string

const (
    ActivityActoinLogin             ActivityActoin = "Login"
    ActivityActoinLogout            ActivityActoin = "Logout"
    ActivityActionUpdate            ActivityActoin = "Update"
    ActivityActionCreate            ActivityActoin = "Add"
    ActivityActionDelete            ActivityActoin = "Delete"
    ActivityActionChangePassword    ActivityActoin = "ChangePassword"
    ActivityActionUpdateLabel       ActivityActoin = "UpdateLabel"
    ActivityActionEnable            ActivityActoin = "Enable"
    ActivityActionDisable           ActivityActoin = "Disable"
    ActivityActionAddNode           ActivityActoin = "AddNode"
    ActivityActionRemoveNode        ActivityActoin = "RemoveNode"
    ActivityActionUpdateBasic       ActivityActoin = "UpdateBasic"
    ActivityActionUpdateApplication ActivityActoin = "UpdateApplication"
    ActivityActionUpdateDeployment  ActivityActoin = "UpdateDeployment"
    ActivityActionStart             ActivityActoin = "Start"
    ActivityActionRestart           ActivityActoin = "Restart"
    ActivityActionStop              ActivityActoin = "Stop"
    ActivityActionStartInstance     ActivityActoin = "StartInstance"
    ActivityActionRestartInstance   ActivityActoin = "RestartInstance "
    ActivityActionStopInstance      ActivityActoin = "StopInstance"
)

type ActivityInfo struct {
    ActivityId   string         `json:"activityId"`
    Action       ActivityActoin `json:"action"`
    Description  string         `json:"description"`
    OldContent   any            `json:"oldContent"`
    NewContent   any            `json:"newContent"`
    Operator     string         `json:"operator"`
    OperatorId   string         `json:"operatorId"`
    OperateAt    int64          `json:"operateAt"`
    ResourceId   string         `json:"resourceId"`
    ResourceName string         `json:"resourceName"`
}
