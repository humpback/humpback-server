package scheduler

import (
	"net/http"
	"time"

	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"

	"github.com/gin-gonic/gin"
)

func getAllNodes(c *gin.Context) {

	nodes, err := db.GetDataAll[types.Node](db.BucketNodes)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nodes)
	}
}

func mockNodes(c *gin.Context) {

	node1 := &types.Node{
		NodeId:    utils.GenerateRandomStringWithLength(8),
		Name:      "hb001",
		IpAddress: "172.30.198.172",
		Port:      8018,
		Status:    "Online",
		IsEnable:  true,
		CreatedAt: time.Now().Unix(),
	}

	db.SaveData(db.BucketNodes, node1.NodeId, node1)

	node2 := &types.Node{
		NodeId:    utils.GenerateRandomStringWithLength(8),
		Name:      "hb002",
		IpAddress: "172.16.41.22",
		Port:      8018,
		Status:    "Online",
		IsEnable:  true,
		CreatedAt: time.Now().Unix(),
	}

	db.SaveData(db.BucketNodes, node2.NodeId, node2)

	node3 := &types.Node{
		NodeId:    utils.GenerateRandomStringWithLength(8),
		Name:      "hb003",
		IpAddress: "172.16.41.23",
		Port:      8018,
		Status:    "Online",
		IsEnable:  true,
		CreatedAt: time.Now().Unix(),
	}

	db.SaveData(db.BucketNodes, node3.NodeId, node3)

	group1 := &types.NodesGroups{
		GroupId:   "GroupTest",
		GroupName: "GroupTest",
		CreatedAt: time.Now().Unix(),
		Nodes:     []string{node1.NodeId, node2.NodeId, node3.NodeId},
	}

	db.SaveData(db.BucketNodesGroups, group1.GroupId, group1)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func getAllServices(c *gin.Context) {

	svcs, err := db.GetDataAll[types.Service](db.BucketServices)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, svcs)
	}
}

func mockGatewayServices(c *gin.Context) {

	svc := &types.Service{
		ServiceId:   utils.GenerateRandomStringWithLength(8),
		ServiceName: "Gateway",
		Version:     utils.GenerateRandomStringWithLength(5),
		IsEnabled:   true,
		Status:      types.ServiceStatusNotReady,
		GroupId:     "GroupTest",
		CreateAt:    time.Now().Unix(),
		Deployment: &types.Deployment{
			Type: types.DeployTypeBackground,
			Mode: types.DeployModeGlobal,
		},
	}

	db.SaveData(db.BucketServices, svc.ServiceId, svc)

	sc := c.MustGet("scheduler").(*HumpbackScheduler)

	svcChange := ServiceChangeInfo{
		ServiceId: svc.ServiceId,
		Version:   svc.Version,
	}

	sc.ServiceChangeChan <- svcChange

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func mockWebServices(c *gin.Context) {

	svc := &types.Service{
		ServiceId:   utils.GenerateRandomStringWithLength(8),
		ServiceName: "Http Web Service",
		Version:     utils.GenerateRandomStringWithLength(5),
		IsEnabled:   true,
		Status:      types.ServiceStatusNotReady,
		GroupId:     "GroupTest",
		CreateAt:    time.Now().Unix(),
		Deployment: &types.Deployment{
			Type:     types.DeployTypeBackground,
			Mode:     types.DeployModeReplicate,
			Replicas: 2,
		},
		Meta: &types.ServiceMetaDocker{
			Image: "nginx:latest",
			Network: &types.NetworkInfo{
				Mode: types.NetworkModeBridge,
				Ports: []*types.PortInfo{
					{
						HostPort:      0,
						ContainerPort: 80,
					},
				},
			},
			RestartPolicy: &types.RestartPolicy{
				Mode: types.RestartPolicyModeAlways,
			},
		},
	}

	db.SaveData(db.BucketServices, svc.ServiceId, svc)

	sc := c.MustGet("scheduler").(*HumpbackScheduler)

	svcChange := ServiceChangeInfo{
		ServiceId: svc.ServiceId,
		Version:   svc.Version,
	}

	sc.ServiceChangeChan <- svcChange

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
