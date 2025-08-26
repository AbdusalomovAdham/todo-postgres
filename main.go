package main

import (
	"myproject/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	routes.UserRouter(r)
	routes.TaskRouter(r)
	r.Run(":8085")
}
