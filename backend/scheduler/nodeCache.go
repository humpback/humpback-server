package scheduler

import (
	"slices"
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

func MatchNodeWithIpAddress(ipAddress []string) string {
	for _, ip := range ipAddress {
		if v, ok := cache.Get(ip); ok {
			return v.(string)
		}
	}

	n, err := db.GetDataByQuery[types.Node](db.BucketNodes, func(key string, value interface{}) bool {
		node := value.(*types.Node)
		return slices.Contains(ipAddress, node.IpAddress)
	})

	id := ""
	if err == nil && len(n) > 0 {
		id = n[0].NodeId
		cache.Add(n[0].IpAddress, id)
	} else {
		cache.Add(ipAddress[0], id)
	}

	return id
}

func GetNodeInfo(nodeId string) *types.Node {
	if v, ok := nodeCache.Get(nodeId); ok {
		return v.(*types.Node)
	}

	n, err := db.NodeGetById(nodeId)
	if err != nil {
		return nil
	}

	nodeCache.Add(nodeId, n)

	return n
}
