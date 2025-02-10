package controller

import (
	"id-information/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: domain.CodeSuccess,
		Data: data,
		Msg:  "success",
	})
}

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}

func ResponseErrorWithApiError(c *gin.Context, apiError *domain.ApiError) {
	if apiError == nil {
		ResponseError(c, 0, "unknown error")
		return
	}
	c.JSON(http.StatusOK, Response{
		Code: apiError.Code,
		Data: nil,
		Msg:  apiError.Message,
	})
}
