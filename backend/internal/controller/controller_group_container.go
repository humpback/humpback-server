package controller

import (
	"sync"
	"time"

	"humpback/api/handle/models"
	"humpback/internal/node"
	"humpback/types"
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

func GroupContainerPerformances(containers models.GroupContainerPerformanceReqInfo) ([]*types.GroupContainerPerformance, error) {
	type tempResult struct {
		NodeId      string
		ContaienrId string
		Stats       *node.ContainerStats
		Err         error
	}
	var (
		l       sync.Mutex
		wg      = &sync.WaitGroup{}
		tempMap = map[string]*tempResult{}
		result  = make([]*types.GroupContainerPerformance, 0)
	)

	for _, info := range containers {
		wg.Add(1)
		go func(info *models.GroupContainerStatsReqInfo) {
			stats, err := node.GetContainerStats(info.NodeId, info.ContainerId)
			l.Lock()
			tempMap[info.ContainerId] = &tempResult{NodeId: info.NodeId, ContaienrId: info.ContainerId, Stats: stats, Err: err}
			l.Unlock()
			wg.Done()
		}(info)
	}
	wg.Wait()
	statsAt := time.Now().UnixMilli()
	for _, p := range tempMap {
		t := &types.GroupContainerPerformance{
			ContainerId: p.ContaienrId,
			NodeId:      p.NodeId,
			IsSuccess:   p.Err == nil,
			Error:       "",
			Stats:       nil,
		}
		if p.Err == nil {
			t.Stats = &types.GroupContainerStats{
				CpuPercent:  p.Stats.CpuPercent,
				MermoryUsed: p.Stats.MermoryUsed,
				MemoryLimit: p.Stats.MemoryLimit,
				IORead:      p.Stats.DiskReadBytes,
				IOWrite:     p.Stats.DiskWriteBytes,
				StatsAt:     statsAt,
				Networks:    p.Stats.Networks,
			}
		} else {
			t.Error = p.Err.Error()
		}
		result = append(result, t)
	}
	return result, nil
}
