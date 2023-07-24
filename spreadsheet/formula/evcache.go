package formula

import "sync"

// evCache is a struct with collection of caching methods intended for add cache support to evaluators.
type evCache struct {
	cache map[string]Result
	lock  *sync.Mutex
}

func newEvCache() evCache {
	ev := evCache{}
	ev.cache = make(map[string]Result)
	ev.lock = &sync.Mutex{}
	return ev
}

func (ec *evCache) SetCache(key string, value Result) {
	ec.lock.Lock()
	ec.cache[key] = value
	ec.lock.Unlock()
}

func (ec *evCache) GetFromCache(key string) (Result, bool) {
	ec.lock.Lock()
	result, found := ec.cache[key]
	ec.lock.Unlock()
	return result, found
}
