package scheduler

import (
	"humpback/internal/db"
	"humpback/types"
	"time"

	tlcache "github.com/JamesYYang/go-ttl-lru"
)

var cache *tlcache.Cache

func NewCacheManager() {
	cache = tlcache.NewLRUWithTTLCache(1000, 60*time.Minute)
}

func MatchNodeWithIpAddress(ipAddress string) string {
	if v, ok := cache.Get(ipAddress); ok {
		return v.(string)
	}

	n, err := db.GetDatabyQuery[types.Node](db.BucketNodes, func(key string, value interface{}) bool {
		node := value.(types.Node)
		return node.IpAddress == ipAddress
	})

	id := ""
	if err == nil || len(n) > 0 {
		id = n[0].NodeID
	}
	cache.Add(ipAddress, id)

	return id
}
