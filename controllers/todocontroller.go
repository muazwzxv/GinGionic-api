package controllers

import (
	"Go-Learn-API/auth"
	"Go-Learn-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTodo :=
func CreateTodo(ctx *gin.Context) {
	var td models.Todo

	if err := ctx.ShouldBindJSON(&td); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	tokenAuth, err := auth.ExtractTokenAuth(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	foundAuth, err := models.Model.FetchAuth(tokenAuth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	td.UserID = foundAuth.UserID
	todo, err := models.Model.CreateTodo(&td)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, todo)
}
