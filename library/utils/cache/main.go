package cache

import (
	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
)

type cache struct {
}

var (
	Cache       = new(cache)
	userRedis   = g.Cfg().GetBool("redis.open")
	GCacheRedis = Cache.New()
)

func (s *cache) New() *gcache.Cache {
	c := gcache.New()
	if userRedis {
		adapter := adapter.NewRedis(g.Redis())
		c.SetAdapter(adapter)
	}

	return c
}
