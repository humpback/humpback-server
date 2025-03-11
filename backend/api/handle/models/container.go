package models

import (
	"slices"
	"strings"

	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/types"
)

type GroupContainerOperateReqInfo struct {
	GroupId     string `json:"-"`
	ContainerId string `json:"containerId"`
	NodeId      string `json:"nodeId"`
	Action      string `json:"action"`
}

func (g *GroupContainerOperateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(g.NodeId, locales.CodeNodesIdNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckIsEmpty(g.ContainerId, locales.CodeContainerIdNotEmpty); err != nil {
		return err
	}
	var actions = []string{strings.ToLower(types.ServiceActionStart), strings.ToLower(types.ServiceActionRestart), strings.ToLower(types.ServiceActionStop)}
	if slices.Index(actions, strings.ToLower(g.Action)) == -1 {
		return response.NewBadRequestErr(locales.CodeContainerActionInvalid)
	}
	return nil
}

type GroupContainerLogsReqInfo struct {
	GroupId       string `json:"-"`
	ContainerId   string `json:"containerId"`
	NodeId        string `json:"nodeId"`
	Line          uint   `json:"line"`
	StartAt       int64  `json:"startAt"`
	EndAt         int64  `json:"endAt"`
	ShowTimestamp bool   `json:"showTimestamp"`
}

func (g *GroupContainerLogsReqInfo) Check() error {
	if err := verify.CheckIsEmpty(g.ContainerId, locales.CodeContainerIdNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckIsEmpty(g.NodeId, locales.CodeNodesIdNotEmpty); err != nil {
		return err
	}
	if g.StartAt > g.EndAt {
		return response.NewBadRequestErr(locales.CodeContainerLogTimeInvlaid)
	}
	if g.Line > 10000 {
		g.Line = 10000
	}
	return nil
}
