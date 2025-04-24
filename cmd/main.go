package main

import (
	"patientsportal/config"
	"patientsportal/models"
	"patientsportal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	models.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
