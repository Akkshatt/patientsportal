package routes

import (
	"patientsportal/controllers"
	"patientsportal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Registration & Login
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Patient CRUD (Protected)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware()) // Or skip this for testing
	{
		authorized.POST("/patients", controllers.CreatePatient)
		authorized.GET("/patients", controllers.GetPatients)
		authorized.GET("/patients/:id", controllers.GetPatientByID)
		authorized.PUT("/patients/:id", controllers.UpdatePatient)
		authorized.DELETE("/patients/:id", controllers.DeletePatient)
	}
}
