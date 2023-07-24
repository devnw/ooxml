package formula

// noCache is a struct with collection of caching methods stubs intended for evaluators without cache.
type noCache struct{}

func (nc *noCache) SetCache(key string, value Result) {}

func (nc *noCache) GetFromCache(key string) (Result, bool) {
	return empty, false
}
