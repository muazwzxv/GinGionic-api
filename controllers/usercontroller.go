package controllers

import (
	"Go-Learn-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser := creates new user
func CreateUser(ctx *gin.Context) {
	var newuser models.User

	if err := ctx.ShouldBindJSON(&newuser); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	user, err := models.Model.CreateUser(&newuser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
