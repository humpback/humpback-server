package controller

import (
	"fmt"
	"log/slog"
	"slices"

	"humpback/config"
	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"
)

type ControllerInter interface {
}

var Controller ControllerInter = &controller{}

type controller struct{}

func Start(stopCh <-chan struct{}) {
	go SessionGCInterval(stopCh)
}

func InitData() error {
	if err := initAdminUser(); err != nil {
		return err
	}
	if err := initRegistry(); err != nil {
		return err
	}
	return nil
}

func initAdminUser() error {
	slog.Info("[Init Supper Admin] Account check start...")
	adminConfig := config.AdminArgs()
	user, err := db.UserGetSupperAdmin()
	if err != nil {
		return fmt.Errorf("Check supper admin account failed: %s", err)
	}
	if user == nil {
		var (
			t  = utils.NewActionTimestamp()
			id = utils.NewGuidStr()
		)
		if err = db.UserUpdate(id, &types.User{
			UserId:    id,
			Username:  adminConfig.Username,
			Email:     "",
			Password:  adminConfig.Password,
			Phone:     "",
			Role:      types.UserRoleSupperAdmin,
			CreatedAt: t,
			UpdatedAt: t,
			Teams:     nil,
		}); err != nil {
			return fmt.Errorf("Create supper admin account failed: %s", err)
		}
	}
	slog.Info("[Init Supper Admin] Account check completed.")
	return nil
}

func initRegistry() error {
	slog.Info("[Init Default Registry] Default Registry check start...")
	registries, err := db.RegistryGetAll()
	if err != nil {
		return fmt.Errorf("Check default registry failed: %s", err)
	}
	if slices.IndexFunc(registries, func(item *types.Registry) bool {
		return item.URL == "docker.io"
	}) == -1 {
		nowT := utils.NewActionTimestamp()
		if err = db.RegistryUpdate([]*types.Registry{
			{
				RegistryId: utils.NewGuidStr(),
				URL:        "docker.io",
				IsDefault:  true,
				Username:   "",
				Password:   "",
				CreatedAt:  nowT,
				UpdatedAt:  nowT,
			},
		}); err != nil {
			return fmt.Errorf("Create default registry failed: %s", err)
		}
	}
	slog.Info("[Init Default Registry] Default Registry check completed.")
	return nil
}

func sendServiceEvent(serviceChangeChan chan types.ServiceChangeInfo, serviceId, version, action string) {
	if serviceChangeChan != nil {
		serviceChangeChan <- types.ServiceChangeInfo{
			Action:    action,
			ServiceId: serviceId,
			Version:   version,
		}
	}
}

func sendNodeEvent(nodeChangeChan chan types.NodeSimpleInfo, nodeId, status string) {
	if nodeChangeChan != nil {
		nodeChangeChan <- types.NodeSimpleInfo{
			NodeId: nodeId,
			Status: status,
		}
	}
}
