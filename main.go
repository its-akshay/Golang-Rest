package main

import (
	"github.com/gin-gonic/gin"
	"res.mod/db"
	"res.mod/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //local host in which we are listeling request

}
