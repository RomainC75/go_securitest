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
	time    time.Time
	content []byte
}

type MemorySaving struct {
	// sync.Map
	mp         map[string]CacheData
	mu         *sync.Mutex
	maxElapsed time.Duration
}

var MAX_ELAPSED time.Duration = time.Second * 5

func NewMemorySaving() *MemorySaving {
	return &MemorySaving{
		mp:         map[string]CacheData{},
		mu:         &sync.Mutex{},
		maxElapsed: MAX_ELAPSED,
	}
}

func (ms *MemorySaving) Set(key string, data []byte) {
	d := CacheData{
		time:    time.Now(),
		content: data,
	}
	ms.mu.Lock()
	ms.mp[key] = d
	ms.mu.Unlock()
}

func (ms *MemorySaving) Get(key string) ([]byte, error) {
	if d, ok := ms.mp[key]; ok {
		// fmt.Println("-> d :", d.time, d.content[:20])
		if time.Since(d.time) < ms.maxElapsed {
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
