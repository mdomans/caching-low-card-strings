package caching_low_card_strings

type SmallCache struct {
	keyValueOffsetMap map[string]int
	valueOffsetMap    map[string]int
	values            []string
}

func NewSmallCache() *SmallCache {
	s := SmallCache{
		keyValueOffsetMap: map[string]int{},
		valueOffsetMap:    map[string]int{},
		values:            make([]string, 10),
	}
	return &s
}

func (cache *SmallCache) Set(key string, value string) {
	if offset, ok := cache.valueOffsetMap[value]; ok {
		cache.keyValueOffsetMap[key] = offset
	} else {
		cache.values = append(cache.values, value)
		offset = len(cache.values)
		cache.keyValueOffsetMap[key] = offset
		cache.valueOffsetMap[value] = offset
	}
}
