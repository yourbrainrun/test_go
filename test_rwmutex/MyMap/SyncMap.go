package MyMap

import (
	"sync"
	"time"
)

type shareData struct {
	MyMap map[string]string
	Lock  sync.RWMutex
}

func NewShareData() *shareData {
	data := shareData{
		MyMap: map[string]string{
			"test": "OK",
		},
		Lock: sync.RWMutex{},
	}

	return &data
}

func (s *shareData) Get(key string, sec int) string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	time.Sleep(time.Duration(sec) * time.Millisecond)
	return s.MyMap[key]
}

func (s *shareData) Set(key string, value string) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.MyMap[key] = value
	return nil
}
