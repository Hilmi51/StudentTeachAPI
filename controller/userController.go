package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-teach-api/config"
	"student-teach-api/models"
)

func ListUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, &users)

}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &user)
	}

}
