package models

import (
	"humpback/common/locales"
	"humpback/common/verify"
	"humpback/types"
)

type TeamCreateReqInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Users       []string `json:"users"`
}

func (t *TeamCreateReqInfo) Check() error {
	if err := verify.CheckRequiredAndLengthLimit(t.Name, locales.LimitTeamName.Min, locales.LimitTeamName.Max, locales.CodeTeamNameNotEmpty, locales.CodeTeamNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(t.Description, 0, locales.LimitDescription.Max, locales.CodeDescriptionLimitLength); err != nil {
		return err
	}
	if len(t.Users) == 0 {
		t.Users = make([]string, 0)
	}
	return nil
}

type TeamUpdateReqInfo struct {
	TeamId string `json:"teamId"`
	TeamCreateReqInfo
}

func (t *TeamUpdateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(t.TeamId, locales.CodeTeamIdNotEmpty); err != nil {
		return err
	}
	return t.TeamCreateReqInfo.Check()
}

type TeamQueryReqInfo struct {
	types.QueryInfo
}

func (t *TeamQueryReqInfo) Check() error {
	t.QueryInfo.CheckBase()
	return nil
}
