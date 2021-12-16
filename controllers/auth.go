package controllers

import (
	"net/http"
	"project/jwt-gin/models"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var  request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	u := models.User{}
	u.Username = request.Username
	u.Password = request.Password
	u.BeforeSave()
	_,err := u.SaveUser()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":err,
		})

}