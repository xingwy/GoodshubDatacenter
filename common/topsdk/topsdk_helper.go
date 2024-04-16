package topsdk_helper

import (
	"context"
	"goodshub-datacenter/conf"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability312"
	request312 "github.com/xingwy/topsdk/ability312/request"
	response312 "github.com/xingwy/topsdk/ability312/response"
	"github.com/xingwy/topsdk/ability648"
	request648 "github.com/xingwy/topsdk/ability648/request"
	response648 "github.com/xingwy/topsdk/ability648/response"
	"github.com/xingwy/topsdk/defaultability"
	"github.com/xingwy/topsdk/defaultability/request"
	"github.com/xingwy/topsdk/defaultability/response"
)

type TopSDKHelper interface {
	// https://open.taobao.com/api.htm?docId=10824&docType=2&scopeId=10980 客户分层-会员平台分层管理
	TaobaoCrmGradeGet(ctx context.Context, userSession string, req *request.TaobaoCrmGradeGetRequest) (*response.TaobaoCrmGradeGetResponse, error)
	// https://open.taobao.com/api.htm?docId=28451&docType=2&scopeId=24284 客户分层-客户运营平台-标签和人群管理
	TaobaoOpencrmCrowdCheck(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=22816&docType=2&scopeId=11295 订单追踪-订单全链路
	TaobaoJdsTradeTracesGet(ctx context.Context, userSession string, req *request.TaobaoJdsTradeTracesGetRequest) (*response.TaobaoJdsTradeTracesGetResponse, error)
	// https://open.taobao.com/api.htm?docId=57630&docType=2&scopeId=23898 订单追踪-简版订单信息查询
	TaobaoTradeSimpleGet(ctx context.Context, userSession string, req *request.TaobaoTradeSimpleGetRequest) (*response.TaobaoTradeSimpleGetResponse, error)
	// https://open.taobao.com/api.htm?docId=55&docType=2&scopeId=12136 订单追踪-评价信息查询
	TaobaoTraderatesGet(ctx context.Context, userSession string, req *request648.TaobaoTraderatesGetRequest) (*response648.TaobaoTraderatesGetResponse, error)
	// https://open.taobao.com/api.htm?docId=51&docType=2&scopeId=11850 订单追踪-退款管理包
	TaobaoRefundsApplyGet(ctx context.Context, userSession string, req *request.TaobaoRefundsApplyGetRequest) (*response.TaobaoRefundsApplyGetResponse, error)
	// https://open.taobao.com/api.htm?docId=54934&docType=2&scopeId=23401 订单追踪-新版物流发货
	AlibabaAscpLogisticsOfflineSend(ctx context.Context, userSession string, req *request.AlibabaAscpLogisticsOfflineSendRequest) (*response.AlibabaAscpLogisticsOfflineSendResponse, error)
	// https://open.taobao.com/api.htm?docId=25521&docType=2&scopeId=11815 短信群发-聚石塔短信服务
	AlibabaAliqinTaSmsNumSend(ctx context.Context, req *request312.AlibabaAliqinTaSmsNumSendRequest) (*response312.AlibabaAliqinTaSmsNumSendResponse, error)
	// https://open.taobao.com/api.htm?docId=57753&docType=2&scopeId=24227 短信群发-客户运营平台-短信发送(CRM专用)
	TaobaoOpencrmSmsSingleSend(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=59184&docType=2&scopeId=24606 短信群发-客户运营平台-数字短信权限包
	TaobaoOpencrmDigitalSmsSingleSend(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=59185&docType=2&scopeId=24605 短信群发-客户运营平台-卡片短信权限包
	TaobaoOpencrmCardSmsCrowdSend(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=57623&docType=2&scopeId=24046 通用-客户运营平台-nick与openid互转
	TaobaoCrmHistoryOuidGet(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=120&docType=2&scopeId=381 通用-系统工具
	TimeGet(ctx context.Context, userSession string, req any) (any, error)
	// https://open.taobao.com/api.htm?docId=18&docType=2&scopeId=382 商品相关-卖家商品查询
	TaobaoItemsOnsaleGet(ctx context.Context, userSession string, req *request.TaobaoItemsOnsaleGetRequest) (*response.TaobaoItemsOnsaleGetResponse, error)
	// https://open.taobao.com/api.htm?docId=24625&docType=2&scopeId=12138 商品相关-商品同步
	TaobaoItemSellerGet(ctx context.Context, userSession string, req *request.TaobaoItemSellerGetRequest) (*response.TaobaoItemSellerGetResponse, error)
	// https://open.taobao.com/api.htm?docId=122&docType=2&scopeId=383 商品相关-商品类目属性
	TaobaoItemcatsGet(ctx context.Context, req *request.TaobaoItemcatsGetRequest) (*response.TaobaoItemcatsGetResponse, error)
}

type TopSDKInstance struct {
	client *topsdk.TopClient
}

func (i *TopSDKInstance) Init() *TopSDKInstance {
	client := topsdk.NewDefaultTopClient(
		conf.Conf.OpenTaoBao.AppKey,
		conf.Conf.OpenTaoBao.AppSecret,
		conf.Conf.OpenTaoBao.ServerUrl,
		conf.Conf.OpenTaoBao.ConnectTimeount,
		conf.Conf.OpenTaoBao.ReadTimeout,
	)

	i.client = &client

	return i
}

func (i *TopSDKInstance) TaobaoCrmGradeGet(ctx context.Context, userSession string, req *request.TaobaoCrmGradeGetRequest) (*response.TaobaoCrmGradeGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoCrmGradeGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoOpencrmCrowdCheck(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TaobaoJdsTradeTracesGet(ctx context.Context, userSession string, req *request.TaobaoJdsTradeTracesGetRequest) (*response.TaobaoJdsTradeTracesGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoJdsTradeTracesGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoTradeSimpleGet(ctx context.Context, userSession string, req *request.TaobaoTradeSimpleGetRequest) (*response.TaobaoTradeSimpleGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoTradeSimpleGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoTraderatesGet(ctx context.Context, userSession string, req *request648.TaobaoTraderatesGetRequest) (*response648.TaobaoTraderatesGetResponse, error) {
	return ability648.NewAbility648(i.client).TaobaoTraderatesGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoRefundsApplyGet(ctx context.Context, userSession string, req *request.TaobaoRefundsApplyGetRequest) (*response.TaobaoRefundsApplyGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoRefundsApplyGet(req, userSession)
}

func (i *TopSDKInstance) AlibabaAscpLogisticsOfflineSend(ctx context.Context, userSession string, req *request.AlibabaAscpLogisticsOfflineSendRequest) (*response.AlibabaAscpLogisticsOfflineSendResponse, error) {
	return defaultability.NewDefaultability(i.client).AlibabaAscpLogisticsOfflineSend(req, userSession)
}

func (i *TopSDKInstance) AlibabaAliqinTaSmsNumSend(ctx context.Context, req *request312.AlibabaAliqinTaSmsNumSendRequest) (*response312.AlibabaAliqinTaSmsNumSendResponse, error) {
	return ability312.NewAbility312(i.client).AlibabaAliqinTaSmsNumSend(req)
}

func (i *TopSDKInstance) TaobaoOpencrmSmsSingleSend(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TaobaoOpencrmDigitalSmsSingleSend(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TaobaoOpencrmCardSmsCrowdSend(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TaobaoCrmHistoryOuidGet(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TimeGet(ctx context.Context, userSession string, req any) (any, error) {
	return nil, nil
}

func (i *TopSDKInstance) TaobaoItemsOnsaleGet(ctx context.Context, userSession string, req *request.TaobaoItemsOnsaleGetRequest) (*response.TaobaoItemsOnsaleGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoItemsOnsaleGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoItemSellerGet(ctx context.Context, userSession string, req *request.TaobaoItemSellerGetRequest) (*response.TaobaoItemSellerGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoItemSellerGet(req, userSession)
}

func (i *TopSDKInstance) TaobaoItemcatsGet(ctx context.Context, req *request.TaobaoItemcatsGetRequest) (*response.TaobaoItemcatsGetResponse, error) {
	return defaultability.NewDefaultability(i.client).TaobaoItemcatsGet(req)
}
