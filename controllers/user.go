package controllers

import (
	"net/http"
	"time"

	"demo/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	IUserService services.IUserService
}

func (_self UserController) Register(ctx *gin.Context) {

	//err := _self.IUserService.Register()

	ctx.JSON(http.StatusOK, gin.H{
		"Time":    time.Now(),
		"TimeUTC": time.Now().UTC(),
		//"Error":   err.Error(),
	})
}
