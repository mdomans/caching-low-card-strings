package caching_low_card_strings

import "github.com/alphadose/haxmap"

type HaxCache struct {
	data *haxmap.Map[string, string]
}

type SmallHaxCache struct {
	keyValueOffsetMap *haxmap.Map[string, int]
	valueOffsetMap    map[string]int
	values            []string
}

func NewHaxCache() *HaxCache {
	s := HaxCache{
		data: haxmap.New[string, string](),
	}
	return &s
}

func (cache *HaxCache) Set(key string, value string) {
	//cache.data[key] = value
	cache.data.Set(key, key)
}

func NewSmallHaxCache() *SmallHaxCache {
	s := SmallHaxCache{
		keyValueOffsetMap: haxmap.New[string, int](),
		valueOffsetMap:    map[string]int{},
		values:            []string{},
	}
	return &s
}

func (cache *SmallHaxCache) Set(key string, value string) {
	if offset, ok := cache.valueOffsetMap[value]; ok {
		cache.keyValueOffsetMap.Set(key, offset)
	} else {
		cache.values = append(cache.values, value)
		offset = len(cache.values)
		cache.keyValueOffsetMap.Set(key, offset)
		cache.valueOffsetMap[value] = offset
	}
}
