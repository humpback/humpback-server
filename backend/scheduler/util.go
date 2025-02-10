package scheduler

import (
	"humpback/pkg/utils"
	"humpback/types"
	"strings"
)

func parseVersionByContainerId(containerId string) string {
	parts := strings.Split(containerId, "-")
	if len(parts) < 4 {
		return ""
	}
	return parts[2]
}

func ParseServiceIdByContainerId(containerId string) string {
	parts := strings.Split(containerId, "-")
	if len(parts) < 4 {
		return ""
	}
	return parts[1]
}

func GenerateContainerName(serviceId, version string) string {
	return "humpback-" + serviceId + "-" + version + "-" + utils.GenerateRandomStringWithLength(5)
}

func isContainerRunning(status string) bool {
	return status == types.ContainerStatusRunning
}

func isContainerExited(status string) bool {
	return status == types.ContainerStatusExited
}

func isContainerRemoved(status string) bool {
	return status == types.ContainerStatusRemoved
}

func isContainerStarting(status string) bool {
	return status == types.ContainerStatusPending ||
		status == types.ContainerStatusStarting ||
		status == types.ContainerStatusCreated
}

func isContainerFailed(status string) bool {
	return status == types.ContainerStatusFailed ||
		status == types.ContainerStatusWarning
}

func isPlacementMatched(node *types.Node, p *types.PlacementInfo) bool {
	if p.Mode == types.PlacementModeIP {
		if p.IsEqual {
			return node.IpAddress == p.Value
		} else {
			return node.IpAddress != p.Value
		}
	} else {
		label := p.Key
		if l, ok := node.Labels[label]; ok {
			if p.IsEqual {
				return l == p.Value
			} else {
				return l != p.Value
			}
		} else {
			return false
		}
	}
}
