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

type LoginRequest struct {
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

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	u := models.User{}
	response := models.LoginResponse{}

	u.Username = request.Username
	u.Password = request.Password

	token,err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = "username or password is incorrect."
		c.JSON(response.Status,gin.H{
			"response":response,
		})
		return
	}

	response.Status = http.StatusOK
	response.Token = token
	c.JSON(http.StatusOK, gin.H{
		"response":response,
	})

}