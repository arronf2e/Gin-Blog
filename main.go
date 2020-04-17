package main

import (
	"gin-blog/config"
	"gin-blog/routers"
)

var err error

func main() {
	config.InitSQL()
	// config.CloseSQL()
	routers.SetupRouter()
}
