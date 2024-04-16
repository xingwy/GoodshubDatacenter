package utils

import (
	"goodshub-datacenter/constants"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Data any    `json:"data"`
	Code int64  `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func ResponseOk(ctx *gin.Context, data any, msg string) {

	ctx.JSON(200, &CommonResponse{
		Code: int64(constants.RESPONSE_TYPE__RESPONSE_OK),
		Msg:  msg,
		Data: data,
	})
}

func ResponseFailed(ctx *gin.Context, msg string) {
	ctx.JSON(500, &CommonResponse{
		Code: int64(constants.RESPONSE_TYPE__RESPONSE_OK),
		Msg:  msg,
		Data: nil,
	})
}
