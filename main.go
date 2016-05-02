package main

import "github.com/vjscjp/sampleutil/core/server"

const (
	Port = "8888"
)

func main() {
	server := server.InitRoutes()
	server.Run(":" + Port)
}
