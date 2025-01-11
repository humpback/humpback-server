package scheduler

import (
	"log"
	"sync/atomic"
	"time"

	"humpback/config"
	"humpback/internal/db"
	"humpback/types"

	"github.com/samber/lo"
	"golang.org/x/exp/rand"
)

type ServiceManager struct {
	ServiceInfo    *types.Service
	CheckInterval  int64
	IsNeedCheckAll atomic.Value
}

func NewServiceManager(svc *types.Service) *ServiceManager {
	sm := &ServiceManager{
		ServiceInfo:   svc,
		CheckInterval: int64(config.BackendArgs().ServiceCheckInterval),
	}

	sm.IsNeedCheckAll.Store(true)

	go sm.CheckService()

	return sm
}

// Reconcile 服务状态变化时，调用该方法，让服务逐步趋于预期状态
func (sm *ServiceManager) Reconcile() {
	if sm.IsNeedCheckAll.Load().(bool) {
		sm.IsNeedCheckAll.Store(false)
		// check node status

		// check service version
	} else {
		// check container status
	}
}

// UpdateContainerWhenChanged 如果容器状态有变化，就保存DB
// 然后等定时检查起来后，重新Reconcile Service
func (sm *ServiceManager) UpdateContainerWhenChanged(cs types.ContainerStatus) {

	ct, ok := lo.Find(sm.ServiceInfo.Containers, func(c *types.ContainerStatus) bool {
		return c.ContainerId == cs.ContainerId
	})

	if ok && (ct.Status != cs.Status || ct.StartAt != cs.StartAt) {
		ct.Status = cs.Status
		ct.StartAt = cs.StartAt
		db.SaveService(sm.ServiceInfo)
	}

	if !ok {
		sm.ServiceInfo.Containers = append(sm.ServiceInfo.Containers, &cs)
		db.SaveService(sm.ServiceInfo)
	}

}

// 定时检查服务状态，看是否满足预期
func (sm *ServiceManager) CheckService() {
	interval := sm.CheckInterval
	time.Sleep(time.Duration(rand.Int31n(int32(interval))) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		log.Printf("check service [%s]......", sm.ServiceInfo.ServiceId)
		sm.Reconcile()
	}
}
