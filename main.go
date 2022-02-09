package main

import (
	"tagus/config"
	"tagus/model"
	"tagus/router"
)

func main() {
	r := router.SetRouter()
	config.Init()
	model.Init()
	r.Run(":8080")
}
