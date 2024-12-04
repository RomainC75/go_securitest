package cache

import (
	"fmt"
	"net/http"
	responsewriter "server/internal/api/cache/responseWriter"
	"time"
)

type Data struct {
	time    time.Time
	content any
}

type Cache struct {
	data     map[string][]byte
	fun      func(http.ResponseWriter, *http.Request)
	strategy ICachStrategy
}

func NewCache(fn func(http.ResponseWriter, *http.Request)) *Cache {
	return &Cache{
		fun: fn,
	}
}

func (c *Cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := (r.Context().Value("user_id")).(int64)
	rawUri := r.URL.RequestURI()

	key := createCacheEndpoint(userId, rawUri)
	if cached, ok := c.data[key]; ok {
		w.Write(cached)

		// TODO delete
		return
	}

	fw := responsewriter.NewResponseWriter(&w)
	c.fun(fw, r)

	b := fw.Get()
	fmt.Println("---> byte to write : ", b)

	fw.Send()
}

func createCacheEndpoint(userId int64, rawUri string) string {
	return fmt.Sprintf("%d-%s", userId, rawUri)
}
