package scheduler

import (
	"log"
	"time"

	"humpback/config"
	"humpback/types"

	"golang.org/x/exp/rand"
)

type ServiceManager struct {
	ServiceInfo   *types.Service
	CheckInterval int64
}

func NewServiceManager(svc *types.Service) *ServiceManager {
	sm := &ServiceManager{
		ServiceInfo:   svc,
		CheckInterval: int64(config.BackendArgs().ServiceCheckInterval),
	}

	return sm
}

func (sm *ServiceManager) Reconcile() {
}

// UpdateContainerWhenChanged 如果容器状态有变化，就保存DB
func (sm *ServiceManager) UpdateContainerWhenChanged(cs types.ContainerStatus) {

}

func (sm *ServiceManager) CheckService() {
	interval := sm.CheckInterval
	time.Sleep(time.Duration(rand.Int31n(int32(interval))) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		log.Printf("check service [%s]......", sm.ServiceInfo.ServiceId)
		sm.CheckServiceCore()
	}
}

func (sm *ServiceManager) CheckServiceCore() {

}
