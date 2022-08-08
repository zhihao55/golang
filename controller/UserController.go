package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

func (c UserController) Index(context *gin.Context)  {
	context.JSON(http.StatusOK,"ssss")
}