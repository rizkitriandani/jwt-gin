package main

import (
	"fmt"
	"project/jwt-gin/controllers"
	"project/jwt-gin/models"
	"github.com/gin-gonic/gin"

)

func main() {
	fmt.Println("Helo")
	models.ConnectDatabase()
	r := gin.Default()
	
	public := r.Group("/api")

	public.POST("/register",controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")
}