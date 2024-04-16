package http

import "github.com/gin-gonic/gin"

// Init 路由
func Reginster(instance *gin.Engine) *gin.Engine {
	api := instance.Group("/sh-goodshub-datacenter")

	taobao := api.Group("/taobao")
	{
		taobao.POST("/taobaoCrmGradeGet", TaobaoCrmGradeGet)
		taobao.POST("/taobaoOpencrmCrowdCheck", TaobaoOpencrmCrowdCheck)
		taobao.POST("/taobaoJdsTradeTracesGet", TaobaoJdsTradeTracesGet)
		taobao.POST("/taobaoTradeSimpleGet", TaobaoTradeSimpleGet)
		taobao.POST("/taobaoTraderatesGet", TaobaoTraderatesGet)
		taobao.POST("/taobaoRefundsApplyGet", TaobaoRefundsApplyGet)
		taobao.POST("/alibabaAscpLogisticsOfflineSend", AlibabaAscpLogisticsOfflineSend)
		taobao.POST("/alibabaAliqinTaSmsNumSend", AlibabaAliqinTaSmsNumSend)
		taobao.POST("/taobaoOpencrmSmsSingleSend", TaobaoOpencrmSmsSingleSend)
		taobao.POST("/taobaoItemsOnsaleGet", TaobaoItemsOnsaleGet)
		taobao.POST("/taobaoItemSellerGet", TaobaoItemSellerGet)
		taobao.POST("/taobaoItemcatsGet", TaobaoItemcatsGet)
	}

	return instance
}
