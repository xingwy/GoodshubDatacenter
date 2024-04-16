package http

import (
	"goodshub-datacenter/api/handler"
	"goodshub-datacenter/utils"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	var (
		params any
	)
	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.DebugHandler().Test(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}
	utils.ResponseOk(ctx, data, "success")
}
