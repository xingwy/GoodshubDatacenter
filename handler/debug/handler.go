package debug

import (
	"context"
	"fmt"
	"goodshub-datacenter/constants"

	"github.com/xingwy/topsdk/defaultability/request"
)

type DebugHandler interface {
	// 测试
	Test(ctx context.Context, params any) (any, error)
}

var session = "6102b21159bae7155e787f9c964fc13464ZZ357bec7ea2b2665024630"

// var refresh_token = "6101c21ba4c4f4c671a672771bfc1197faZZab03878b23c2665024630"

func (h *DebugInstance) Test(ctx context.Context, params any) (any, error) {
	data, err := h.TopSDKHelper.TaobaoJdsTradeTracesGet(ctx, session, &request.TaobaoJdsTradeTracesGetRequest{
		ReturnUserStatus: &constants.EMPTY_TRUE,
	})

	fmt.Println(data, err)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *DebugInstance) TaobaoCrmGradeGet(ctx context.Context, params any) (any, error) {
	data, err := h.TopSDKHelper.TaobaoCrmGradeGet(ctx, session, &request.TaobaoCrmGradeGetRequest{})

	fmt.Println(data, err)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *DebugInstance) TaobaoJdsTradeTracesGet(ctx context.Context, params any) (any, error) {
	data, err := h.TopSDKHelper.TaobaoJdsTradeTracesGet(ctx, session, &request.TaobaoJdsTradeTracesGetRequest{
		ReturnUserStatus: &constants.EMPTY_TRUE,
	})

	fmt.Println(data, err)
	if err != nil {
		return nil, err
	}

	return data, nil
}
