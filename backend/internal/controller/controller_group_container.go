package controller

import (
	"humpback/api/handle/models"
	"humpback/internal/node"
)

func GroupContainerOperate(info *models.GroupContainerOperateReqInfo) error {
	if err := node.OperateNodeContainer(info.NodeId, info.ContainerId, info.Action); err != nil {
		return err
	}
	return nil
}

func GroupContainerQueryLogs(info *models.GroupContainerLogsReqInfo) (string, error) {
	return "", nil
}
