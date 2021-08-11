package singleton

import (
	"sync"
)

type singleton struct{}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance() *singleton {
	// 使用nil判断节省资源
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
		})
	}
	return instance
}
