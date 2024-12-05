package cache

import (
	"fmt"
	"net/http"
	responsewriter "server/internal/api/cache/responseWriter"
	"server/internal/api/cache/strategies"
)

type ICacheStrategy interface {
	Set(key string, data []byte)
	Get(key string) ([]byte, error)
	DisplayAll()
}

type Cache struct {
	data     map[string][]byte
	fun      func(http.ResponseWriter, *http.Request)
	strategy ICacheStrategy
}

func NewCache(fn func(http.ResponseWriter, *http.Request)) *Cache {
	return &Cache{
		fun:      fn,
		strategy: strategies.GetMemorySaving(),
	}
}

func (c *Cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := (r.Context().Value("user_id")).(int64)
	rawUri := r.URL.RequestURI()

	key := createCacheEndpoint(userId, rawUri)

	// in cache // no error
	if b, err := c.strategy.Get(key); err == nil {
		fmt.Println("---> In Cache ! ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	fw := responsewriter.NewResponseWriter(&w)

	/// run
	c.fun(fw, r)
	_, b := fw.Get()

	c.strategy.Set(key, b)
	c.strategy.DisplayAll()

	fw.Send()
}

func createCacheEndpoint(userId int64, rawUri string) string {
	return fmt.Sprintf("%d-%s", userId, rawUri)
}
