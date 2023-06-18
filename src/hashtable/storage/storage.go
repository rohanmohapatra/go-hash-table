package storage

var store = make(map[string]uint64)

func Get(key string) (int32, uint64) {
	val, ok := store[key]
	if ok {
		return 0, val
	} else {
		return -1, 0
	}
}

func Put(key string, value uint64) {
	store[key] = value
}
