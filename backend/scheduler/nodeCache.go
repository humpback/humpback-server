package scheduler

import (
	"time"

	"humpback/internal/db"
	"humpback/types"

	tlcache "github.com/JamesYYang/go-ttl-lru"
)

var cache *tlcache.Cache
var nodeCache *tlcache.Cache

func NewCacheManager() {
	cache = tlcache.NewLRUWithTTLCache(1000, 60*time.Minute)
	nodeCache = tlcache.NewLRUWithTTLCache(1000, 60*time.Minute)
}

func MatchNodeWithIpAddress(ipAddress string) string {
	if v, ok := cache.Get(ipAddress); ok {
		return v.(string)
	}

	n, err := db.GetDataByQuery[types.Node](db.BucketNodes, func(key string, value interface{}) bool {
		node := value.(types.Node)
		return node.IpAddress == ipAddress
	})

	id := ""
	if err == nil || len(n) > 0 {
		id = n[0].NodeId
	}
	cache.Add(ipAddress, id)

	return id
}

func GetNodeInfo(nodeId string) *types.Node {
	if v, ok := cache.Get(nodeId); ok {
		return v.(*types.Node)
	}

	n, err := db.GetNodeById(nodeId)
	if err != nil {
		return nil
	}

	cache.Add(nodeId, n)

	return n
}
