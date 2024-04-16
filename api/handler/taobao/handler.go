package taobao_handler

import (
	"context"
	"goodshub-datacenter/model/vo"

	"github.com/xingwy/topsdk/defaultability/request"
)

type TaobaoHandler interface {
	TaobaoCrmGradeGet(ctx context.Context, params *vo.TaobaoCrmGradeGetRequest) (any, error)
	TaobaoOpencrmCrowdCheck(ctx context.Context, params *vo.TaobaoOpencrmCrowdCheckRequest) (any, error)
	TaobaoJdsTradeTracesGet(ctx context.Context, params *vo.TaobaoJdsTradeTracesGetRequest) (any, error)
	TaobaoTradeSimpleGet(ctx context.Context, params *vo.TaobaoTradeSimpleGetRequest) (any, error)
	TaobaoTraderatesGet(ctx context.Context, params *vo.TaobaoTraderatesGetRequest) (any, error)
	TaobaoRefundsApplyGet(ctx context.Context, params *vo.TaobaoRefundsApplyGetRequest) (any, error)
	AlibabaAscpLogisticsOfflineSend(ctx context.Context, params *vo.AlibabaAscpLogisticsOfflineSendRequest) (any, error)
	AlibabaAliqinTaSmsNumSend(ctx context.Context, params *vo.AlibabaAliqinTaSmsNumSendRequest) (any, error)
	TaobaoOpencrmSmsSingleSend(ctx context.Context, params *vo.TaobaoOpencrmSmsSingleSendRequest) (any, error)
	TaobaoItemsOnsaleGet(ctx context.Context, params *vo.TaobaoItemsOnsaleGetRequest) (any, error)
	TaobaoItemSellerGet(ctx context.Context, params *vo.TaobaoItemSellerGetRequest) (any, error)
	TaobaoItemcatsGet(ctx context.Context, params *vo.TaobaoItemcatsGetRequest) (any, error)
	TimeGet(ctx context.Context, params *vo.TimeGetRequest) (*vo.TimeGetResponse, error)
	TaobaoCrmHistoryOuidGet(ctx context.Context, params *vo.TaobaoCrmHistoryOuidGetRequest) (any, error)
}

func (h *TaobaoInstance) TaobaoCrmGradeGet(ctx context.Context, params *vo.TaobaoCrmGradeGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoCrmGradeGet(ctx, params.SessionID, &request.TaobaoCrmGradeGetRequest{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoOpencrmCrowdCheck(ctx context.Context, params *vo.TaobaoOpencrmCrowdCheckRequest) (any, error) {
	return nil, nil
}

func (h *TaobaoInstance) TaobaoJdsTradeTracesGet(ctx context.Context, params *vo.TaobaoJdsTradeTracesGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoJdsTradeTracesGet(ctx, params.SessionID, &params.TaobaoJdsTradeTracesGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoTradeSimpleGet(ctx context.Context, params *vo.TaobaoTradeSimpleGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoTradeSimpleGet(ctx, params.SessionID, &params.TaobaoTradeSimpleGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoTraderatesGet(ctx context.Context, params *vo.TaobaoTraderatesGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoTraderatesGet(ctx, params.SessionID, &params.TaobaoTraderatesGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoRefundsApplyGet(ctx context.Context, params *vo.TaobaoRefundsApplyGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoRefundsApplyGet(ctx, params.SessionID, &params.TaobaoRefundsApplyGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) AlibabaAscpLogisticsOfflineSend(ctx context.Context, params *vo.AlibabaAscpLogisticsOfflineSendRequest) (any, error) {
	data, err := h.TopSDKHelper.AlibabaAscpLogisticsOfflineSend(ctx, params.SessionID, &params.AlibabaAscpLogisticsOfflineSendRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) AlibabaAliqinTaSmsNumSend(ctx context.Context, params *vo.AlibabaAliqinTaSmsNumSendRequest) (any, error) {
	data, err := h.TopSDKHelper.AlibabaAliqinTaSmsNumSend(ctx, &params.AlibabaAliqinTaSmsNumSendRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoOpencrmSmsSingleSend(ctx context.Context, params *vo.TaobaoOpencrmSmsSingleSendRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoOpencrmSmsSingleSend(ctx, params.SessionID, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoItemsOnsaleGet(ctx context.Context, params *vo.TaobaoItemsOnsaleGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoItemsOnsaleGet(ctx, params.SessionID, &params.TaobaoItemsOnsaleGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoItemSellerGet(ctx context.Context, params *vo.TaobaoItemSellerGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoItemSellerGet(ctx, params.SessionID, &params.TaobaoItemSellerGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TaobaoItemcatsGet(ctx context.Context, params *vo.TaobaoItemcatsGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoItemcatsGet(ctx, &params.TaobaoItemcatsGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *TaobaoInstance) TimeGet(ctx context.Context, params *vo.TimeGetRequest) (*vo.TimeGetResponse, error) {
	res := &vo.TimeGetResponse{}
	data, err := h.TopSDKHelper.TimeGet(ctx, &params.TimeGetRequest)
	if err != nil {
		return nil, err
	}

	res.Time = data.Time

	return res, nil
}

func (h *TaobaoInstance) TaobaoCrmHistoryOuidGet(ctx context.Context, params *vo.TaobaoCrmHistoryOuidGetRequest) (any, error) {
	data, err := h.TopSDKHelper.TaobaoCrmHistoryOuidGet(ctx, params.SessionID, &params.TaobaoCrmHistoryOuidGetRequest)
	if err != nil {
		return nil, err
	}

	return data, nil
}
