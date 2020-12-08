package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseHandler struct {
}

func (BaseHandler) OK(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, &APIResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}