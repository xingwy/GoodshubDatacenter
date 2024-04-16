package vo

import (
	"time"

	request312 "github.com/xingwy/topsdk/ability312/request"
	request648 "github.com/xingwy/topsdk/ability648/request"
	"github.com/xingwy/topsdk/defaultability/request"
)

type TaobaoCrmGradeGetRequest struct {
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoCrmGradeGetResponse struct {
}

type TaobaoOpencrmCrowdCheckRequest struct {
}

type TaobaoOpencrmCrowdCheckResponse struct {
}

type TaobaoJdsTradeTracesGetRequest struct {
	request.TaobaoJdsTradeTracesGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoJdsTradeTracesGetResponse struct {
}

type TaobaoTradeSimpleGetRequest struct {
	request.TaobaoTradeSimpleGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoTradeSimpleGetResponse struct {
}

type TaobaoTraderatesGetRequest struct {
	request648.TaobaoTraderatesGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoTraderatesGetResponse struct {
}

type TaobaoRefundsApplyGetRequest struct {
	request.TaobaoRefundsApplyGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoRefundsApplyGetResponse struct {
}

type AlibabaAscpLogisticsOfflineSendRequest struct {
	request.AlibabaAscpLogisticsOfflineSendRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type AlibabaAscpLogisticsOfflineSendResponse struct {
}

type AlibabaAliqinTaSmsNumSendRequest struct {
	request312.AlibabaAliqinTaSmsNumSendRequest
}

type AlibabaAliqinTaSmsNumSendResponse struct {
}

type TaobaoOpencrmSmsSingleSendRequest struct {
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoOpencrmSmsSingleSendResponse struct {
}

type TaobaoItemsOnsaleGetRequest struct {
	request.TaobaoItemsOnsaleGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoItemsOnsaleGetResponse struct {
}

type TaobaoItemSellerGetRequest struct {
	request.TaobaoItemSellerGetRequest
	SessionID string `json:"session_id" form:"session_id"`
}

type TaobaoItemSellerGetResponse struct {
}

type TaobaoItemcatsGetRequest struct {
	request.TaobaoItemcatsGetRequest
}

type TaobaoItemcatsGetResponse struct {
}

type TimeGetRequest struct {
	request.TimeGetRequest
}

type TimeGetResponse struct {
	Time time.Time `json:"time"`
}
