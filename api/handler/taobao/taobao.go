package taobao_handler

import (
	topsdk_helper "goodshub-datacenter/common/topsdk"
	"goodshub-datacenter/dao"
	"goodshub-datacenter/struct/define"
	"goodshub-datacenter/struct/event"
)

type TaobaoInstance struct {
	define.Instance
	// 可以在这里添加类型的属性
	Dao          *dao.Dao
	EventManager event.EventManager
	TopSDKHelper topsdk_helper.TopSDKHelper
}
