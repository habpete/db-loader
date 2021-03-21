package cache

import "github.com/bradfitz/gomemcache/memcache"

type cache struct {
	cache *memcache.Client
}

func New() ICache {
	mc := memcache.New()
	return &cache{
		cache: mc,
	}
}

func (this *cache) Put() {}

func (this *cache) Get() {}
