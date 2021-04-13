package controller

import (
	"github.com/gin-gonic/gin"
	"nspyf/model/dto"
	"nspyf/service"
	"strconv"
)

func PutUserInfo(c *gin.Context) {
	req := &dto.UserInfo{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	data, code := service.PutUserInfo(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func GetUserInfo(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.CommitDataError)
		return
	}

	data, code := service.GetUserInfo(uint(idInt))
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}