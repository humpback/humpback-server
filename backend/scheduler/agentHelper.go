package scheduler

import (
	"fmt"
	"humpback/pkg/httpx"
	"humpback/pkg/utils"
	"humpback/types"
	"log/slog"
)

func RemoveNodeContainer(nodeId string, containerName string) error {
	// remove container
	node := GetNodeInfo(nodeId)
	if node != nil {
		url := fmt.Sprintf("http://%s:%d/containers/%s", node.IpAddress, node.Port, containerName)
		err := httpx.NewHttpXClient().Delete(url, nil, nil, nil)
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
		url := fmt.Sprintf("http://%s:%d/containers", node.IpAddress, node.Port)
		err := httpx.NewHttpXClient().Post(url, nil, nil, svc, nil)
		if err != nil {
			slog.Error("[Agent Helper] Start container error", "error", err.Error())
			return err
		}

		c := &types.ContainerStatus{
			ContainerName: fmt.Sprintf("humpback-%s-%s-%s", svc.ServiceId, svc.Version, utils.GenerateRandomStringWithLength(5)),
			Status:        types.ContainerStatusPending,
			NodeId:        nodeId,
		}
		svc.Containers = append(svc.Containers, c)
	}

	return nil
}
