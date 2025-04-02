package dynamicconfig

import "sync"

var (
	optionStore = make(map[string]string)
	rwMutex     sync.RWMutex
)

// Set 设置动态配置项
func Set(key, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	optionStore[key] = value
}

// Get 获取动态配置项
func Get(key string) (string, bool) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	val, ok := optionStore[key]
	return val, ok
}

// GetAll 返回所有动态配置项的副本
func GetAll() map[string]string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	copyMap := make(map[string]string, len(optionStore))
	for k, v := range optionStore {
		copyMap[k] = v
	}
	return copyMap
}
