package cache

import (
	"fmt"
	"net/http"
)

type Cache struct {
	data map[string][]byte
	fun  func(http.ResponseWriter, *http.Request)
}

func NewCache(fn func(http.ResponseWriter, *http.Request)) *Cache {
	return &Cache{
		fun: fn,
	}
}

func (c *Cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Decorator : ", r.URL)
	c.fun(w, r)
}
