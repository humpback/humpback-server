package scheduler

import (
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

func parseServiceIdByContainerId(containerId string) string {
	parts := strings.Split(containerId, "-")
	if len(parts) < 4 {
		return ""
	}
	return parts[1]
}

func isContainerRunning(status string) bool {
	return status == types.ContainerStatusRunning
}

func isContainerExited(status string) bool {
	return status == types.ContainerStatusExited
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
