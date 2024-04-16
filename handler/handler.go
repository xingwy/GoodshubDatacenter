package handler

import (
	topsdk_helper "goodshub-datacenter/common/topsdk"
	"goodshub-datacenter/conf"
	"goodshub-datacenter/dao"
	"goodshub-datacenter/handler/debug"
	taobao_handler "goodshub-datacenter/handler/taobao"
	"goodshub-datacenter/struct/event"
)

var (
	Instance *Handler
)

type Handler struct {
	dao *dao.Dao
	// 业务 处理器
	taobaoInstance *taobao_handler.TaobaoInstance
	debugInstance  *debug.DebugInstance
}

func New(cfg *conf.Config) (s *Handler) {
	// 数据操作实例
	daoInstance, err := dao.NewDao(cfg)
	if err != nil {
		panic(err)
	}

	// 事件管理器
	eventInstance := event.NewEventManager()
	// 淘宝SDK
	topsdkInstance := (&topsdk_helper.TopSDKInstance{}).Init()

	Instance = &Handler{
		dao: daoInstance,
		debugInstance: &debug.DebugInstance{
			Dao:          daoInstance,
			EventManager: eventInstance,
			TopSDKHelper: topsdkInstance,
		},
		taobaoInstance: &taobao_handler.TaobaoInstance{
			Dao:          daoInstance,
			EventManager: eventInstance,
			TopSDKHelper: topsdkInstance,
		},
	}

	return Instance.Init()
}

func (s *Handler) Init() *Handler {

	return s
}

func (s *Handler) Dao() *dao.Dao {
	return s.dao
}

// 获取 TaobaoHandler
func (s *Handler) TaobaoHandler() taobao_handler.TaobaoHandler {
	return s.taobaoInstance
}

// 获取 DemoHandler
func (s *Handler) DebugHandler() debug.DebugHandler {
	return s.debugInstance
}
