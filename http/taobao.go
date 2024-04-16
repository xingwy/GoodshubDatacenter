package http

import (
	"goodshub-datacenter/handler"
	"goodshub-datacenter/model/vo"
	"goodshub-datacenter/utils"

	"github.com/gin-gonic/gin"
)

func TaobaoCrmGradeGet(ctx *gin.Context) {
	var (
		params vo.TaobaoCrmGradeGetRequest
	)

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoCrmGradeGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoOpencrmCrowdCheck(ctx *gin.Context) {
	var params vo.TaobaoOpencrmCrowdCheckRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoOpencrmCrowdCheck(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoJdsTradeTracesGet(ctx *gin.Context) {
	var params vo.TaobaoJdsTradeTracesGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoJdsTradeTracesGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoTradeSimpleGet(ctx *gin.Context) {
	var params vo.TaobaoTradeSimpleGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoTradeSimpleGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoTraderatesGet(ctx *gin.Context) {
	var params vo.TaobaoTraderatesGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoTraderatesGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoRefundsApplyGet(ctx *gin.Context) {
	var params vo.TaobaoRefundsApplyGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoRefundsApplyGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func AlibabaAscpLogisticsOfflineSend(ctx *gin.Context) {
	var params vo.AlibabaAscpLogisticsOfflineSendRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().AlibabaAscpLogisticsOfflineSend(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func AlibabaAliqinTaSmsNumSend(ctx *gin.Context) {
	var params vo.AlibabaAliqinTaSmsNumSendRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().AlibabaAliqinTaSmsNumSend(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoOpencrmSmsSingleSend(ctx *gin.Context) {
	var params vo.TaobaoOpencrmSmsSingleSendRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoOpencrmSmsSingleSend(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoItemsOnsaleGet(ctx *gin.Context) {
	var params vo.TaobaoItemsOnsaleGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoItemsOnsaleGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoItemSellerGet(ctx *gin.Context) {
	var params vo.TaobaoItemSellerGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoItemSellerGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TaobaoItemcatsGet(ctx *gin.Context) {
	var params vo.TaobaoItemcatsGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TaobaoItemcatsGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}

func TimeGet(ctx *gin.Context) {
	var params vo.TimeGetRequest

	err := ctx.Bind(&params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	data, err := handler.Instance.TaobaoHandler().TimeGet(ctx, &params)
	if err != nil {
		utils.ResponseFailed(ctx, err.Error())
		return
	}

	utils.ResponseOk(ctx, data, "success")
}
