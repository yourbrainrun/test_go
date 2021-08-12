package configs

import (
	"os"
	"sync"
)

type SLSConfig struct {
	AppKey      string
	AppSecret   string
	EndPoint    string
	ProjectName string
	LogStore    string
}

var slsConfigInstance *SLSConfig
var slsConfigOnce sync.Once

func GetSlsConfig() *SLSConfig {
	slsConfigOnce.Do(func() {
		slsConfigInstance = &SLSConfig{
			AppKey:      os.Getenv("ALIYUN_APP_KEY"),
			AppSecret:   os.Getenv("ALIYUN_APP_SECRET"),
			EndPoint:    os.Getenv("ALIYUN_SLS_END_POINT"),
			ProjectName: os.Getenv("ALIYUN_SLS_PROJECT_NAME"),
			LogStore:    os.Getenv("ALIYUN_SLS_STORE_NAME"),
		}
	})
	return slsConfigInstance
}
