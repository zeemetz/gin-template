package main

import (
	"os"
	"template/controller"
	"template/engine"
)

func init() {
	engine.GetAdmin().GET("/ping", controller.Ping)
}

func main() {
	if len(os.Args) < 3 {
		engine.GetEngine().Run("localhost:1313")
	}
	engine.GetEngine().Run(os.Args[1] + ":" + os.Args[2])
}
