package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Manga *cache.Cache

func Init() {
	Manga = cache.New(5*time.Minute, 10*time.Minute)
}
