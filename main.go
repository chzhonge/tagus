package main

import (
	"tagus/cache"
	"tagus/config"
	"tagus/model"
	"tagus/router"
)

func main() {
	r := router.SetRouter()
	config.Init()
	model.Init()
	cache.Init()
	r.Run(":8080")
}
