package scheduler

import (
	"fmt"
	"humpback/pkg/httpx"
	"humpback/pkg/utils"
	"humpback/types"
	"log/slog"
	"strings"
)

func RemoveNodeContainer(nodeId string, containerId string) error {
	// remove container
	node := GetNodeInfo(nodeId)
	if node != nil {
		url := fmt.Sprintf("http://%s:%d/api/v1/container/%s?force=true", node.IpAddress, node.Port, containerId)
		slog.Info("[Agent Helper] Remove container", "url", url)
		err := httpx.NewHttpXClient().Delete(url, nil, nil, nil)
		if err != nil {
			slog.Error("[Agent Helper] Remove container error", "error", err.Error())
			return err
		}
	}
	return nil
}

func OperateNodeContainer(nodeId string, containerId string, action string) error {
	// operate container
	node := GetNodeInfo(nodeId)
	if node != nil {
		url := fmt.Sprintf("http://%s:%d/api/v1/container/%s/%s", node.IpAddress, node.Port, containerId, strings.ToLower(action))
		slog.Info("[Agent Helper] Operate container", "url", url)
		err := httpx.NewHttpXClient().Post(url, nil, nil, nil, nil)
		if err != nil {
			slog.Error("[Agent Helper] Remove container error", "error", err.Error())
			return err
		}
	}
	return nil
}

func StartNewContainer(nodeId string, svc *types.Service) error {
	node := GetNodeInfo(nodeId)
	if node != nil {

		task := &types.AgentTask{
			ContainerName:     GenerateContainerName(svc.ServiceId, svc.Version),
			ServiceMetaDocker: svc.Meta,
			ScheduleInfo:      svc.Deployment.Schedule,
		}
		utils.PrintJson(task)
		url := fmt.Sprintf("http://%s:%d/api/v1/container", node.IpAddress, node.Port)
		slog.Info("[Agent Helper] Create container", "url", url)
		err := httpx.NewHttpXClient().Post(url, nil, nil, task, nil)
		if err != nil {
			slog.Error("[Agent Helper] Start container error", "error", err.Error())
			return err
		}

		c := &types.ContainerStatus{
			ContainerName: task.ContainerName,
			Status:        types.ContainerStatusPending,
			NodeId:        nodeId,
		}
		svc.Containers = append(svc.Containers, c)
	}

	return nil
}
