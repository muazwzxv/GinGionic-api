package controllers

import (
	"Go-Learn-API/auth"
	"Go-Learn-API/models"
	"Go-Learn-API/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login := login user
func Login(ctx *gin.Context) {
	var tologin models.User

	if err := ctx.ShouldBindJSON(&tologin); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Checks if the user exists
	user, err := models.Model.GetUserByEmail(tologin.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	// create the auth record for login
	authdata, err := models.Model.CreateAuth(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var authdetails auth.AuthDetails
	authdetails.UserID = authdata.UserID
	authdetails.AuthUUID = authdata.AuthUUID

	token, loginErr := service.Authorize.SignIn(authdetails)
	if loginErr != nil {
		ctx.JSON(http.StatusForbidden, "please try to login later")
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func logout(ctx *gin.Context) {
	authentcate, err := auth.ExtractTokenAuth(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := models.Model.DeleteAuth(authentcate)
	if delErr != nil {
		log.Println(delErr)
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	ctx.JSON(http.StatusOK, "Successfully logut")
}
