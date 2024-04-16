package main

import (
	"flag"
	"fmt"
	"goodshub-datacenter/conf"
	"goodshub-datacenter/handler"
	"goodshub-datacenter/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
)

type program struct {
	handler    *handler.Handler
	httpEngine *gin.Engine
}

func (p *program) BeforeStart() error {
	fmt.Println("Before Start Program")
	// 初始化处理模块
	p.handler = handler.New(conf.Conf)
	// 注册路由
	p.httpEngine = http.Reginster(gin.Default())
	return nil
}

func (p *program) AfterStart() error {
	fmt.Println("After Start Program")
	// 初始化处理模块
	p.handler = handler.New(conf.Conf)
	// 注册路由
	p.httpEngine = http.Reginster(gin.Default())
	p.httpEngine.Run(conf.Conf.GinServer.Addr())
	return nil
}

func (p *program) BeforeStop() error {
	fmt.Println("Before Stop Program")

	return nil
}

func (p *program) AfterStop() error {
	fmt.Println("After Stop Program")

	return nil
}

func (p *program) Run() error {
	// 运行
	err := p.httpEngine.Run(conf.Conf.GinServer.Addr())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	prg := &program{}
	service := micro.NewService()
	// 初始化
	service.Init(
		micro.Name(conf.Conf.ServiceName),
		micro.BeforeStart(prg.BeforeStart),
		micro.AfterStart(prg.AfterStart),
		micro.BeforeStop(prg.BeforeStop),
		micro.AfterStop(prg.AfterStop),
	)

	if err := service.Run(); err != nil {
		panic(err)
	}
}
