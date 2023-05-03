package caching_low_card_strings

type NaiveCache struct {
	data map[string]string
}

func NewNaiveCache() *NaiveCache {
	s := NaiveCache{
		data: map[string]string{},
	}
	return &s
}

func (cache *NaiveCache) Set(key string, value string) {
	cache.data[key] = value
}
