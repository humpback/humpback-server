package scheduler

import (
	"fmt"
	"humpback/pkg/utils"
	"humpback/types"
)

func RemoveNodeContainer(nodeId string, containerId string) error {
	// remove container
	return nil
}

func StartNewContainer(nodeId string, svc *types.Service) error {
	c := &types.ContainerStatus{
		ContainerName: fmt.Sprintf("humpback-%s-%s-%s", svc.ServiceId, svc.Version, utils.GenerateRandomStringWithLength(5)),
		Status:        types.ContainerStatusPending,
		NodeId:        nodeId,
	}
	svc.Containers = append(svc.Containers, c)
	return nil
}
