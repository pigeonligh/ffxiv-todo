package wiki

import "github.com/pigeonligh/ffxiv-todo/pkg/types"

type Cache interface {
	Get(types.CacheKey) (interface{}, bool)
	Add(types.CacheKey, interface{})
}
