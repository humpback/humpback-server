package scheduler

import (
	"encoding/json"
	"fmt"
	"humpback/pkg/utils"
	"humpback/types"
	"log/slog"
	"strings"
)

func RemoveNodeContainer(nodeId string, containerName string) error {
	// remove container
	node := GetNodeInfo(nodeId)
	if node != nil {
		url := fmt.Sprintf("http://%s:%d/containers/%s", node.IpAddress, node.Port, containerName)
		slog.Info("[Agent Helper] Remove container for mock", "url", url)
		// err := httpx.NewHttpXClient().Delete(url, nil, nil, nil)
		// if err != nil {
		// 	slog.Error("[Agent Helper] Remove container error", "error", err.Error())
		// 	return err
		// }
	}
	return nil
}

func OperateNodeContainer(nodeId string, containerName string, action string) error {
	// remove container
	node := GetNodeInfo(nodeId)
	if node != nil {
		url := fmt.Sprintf("http://%s:%d/containers/%s/%s", node.IpAddress, node.Port, containerName, strings.ToLower(action))
		slog.Info("[Agent Helper] Remove container for mock", "url", url)
		// err := httpx.NewHttpXClient().Delete(url, nil, nil, nil)
		// if err != nil {
		// 	slog.Error("[Agent Helper] Remove container error", "error", err.Error())
		// 	return err
		// }
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

		taskJson, errj := json.Marshal(task)
		if errj != nil {
			slog.Error("[Agent Helper] JSON marshal error", "error", errj.Error())
			return errj
		}
		fmt.Println(string(taskJson))

		url := fmt.Sprintf("http://%s:%d/containers", node.IpAddress, node.Port)
		slog.Info("[Agent Helper] Create container for mock", "url", url)
		// err := httpx.NewHttpXClient().Post(url, nil, nil, task, nil)
		// if err != nil {
		// 	slog.Error("[Agent Helper] Start container error", "error", err.Error())
		// 	return err
		// }

		c := &types.ContainerStatus{
			ContainerName: fmt.Sprintf("humpback-%s-%s-%s", svc.ServiceId, svc.Version, utils.GenerateRandomStringWithLength(5)),
			Status:        types.ContainerStatusPending,
			NodeId:        nodeId,
		}
		svc.Containers = append(svc.Containers, c)
	}

	return nil
}
