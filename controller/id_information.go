package controller

import (
	"id-information/usecase"

	"github.com/gin-gonic/gin"
)

// GetIDInformation 获取身份证信息
func GetIDInformation(c *gin.Context) {
	idNumber := c.Query("id_number")
	resp, err := usecase.ParseIDNumber(idNumber)
	if err != nil {
		ResponseErrorWithApiError(c, err)
		return
	}
	ResponseSuccess(c, resp)
	return
}
