package strategies

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var memorySaving *MemorySaving

func SetMemorySaving() {
	memorySaving = NewMemorySaving()
}

func GetMemorySaving() *MemorySaving {
	return memorySaving
}

type CacheData struct {
	time     time.Time
	content  []byte
	duration time.Duration
}

type MemorySaving struct {
	// sync.Map
	mp map[string]CacheData
	mu *sync.Mutex
}

func NewMemorySaving() *MemorySaving {
	return &MemorySaving{
		mp: map[string]CacheData{},
		mu: &sync.Mutex{},
	}
}

func (ms *MemorySaving) Set(key string, data []byte, duration time.Duration) {
	d := CacheData{
		time:     time.Now(),
		content:  data,
		duration: duration,
	}
	ms.mu.Lock()
	ms.mp[key] = d
	ms.mu.Unlock()
}

func (ms *MemorySaving) Get(key string) ([]byte, error) {
	if d, ok := ms.mp[key]; ok {
		if time.Since(d.time) < d.duration {
			return d.content, nil
		} else {
			ms.Delete(key)
			return []byte{}, errors.New("outdated")
		}
	}
	return []byte{}, errors.New("not found")
}

func (ms *MemorySaving) DisplayAll() {
	// fmt.Println("----DisplayAll()----")
	for k, c := range ms.mp {
		fmt.Println("------> ", k, c.time)
	}
}

func (ms *MemorySaving) Delete(key string) error {
	if _, ok := ms.mp[key]; ok {
		ms.mu.Lock()
		delete(ms.mp, key)
		ms.mu.Unlock()
	}
	return errors.New("not found")
}
