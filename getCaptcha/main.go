package main

import (
	"getCaptcha/handler"
	getCaptcha "getCaptcha/proto/getCaptcha"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

func main() {
	// New Service
	etcd := etcd.NewRegistry(
		// 地址是我本地etcd服务器地址，不要照抄
		registry.Addrs("localhost:2379"),
	)
	service := micro.NewService(
		micro.Name("go.micro.service.getCaptcha"),
		micro.Version("latest"),
		micro.Registry(etcd),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.getCaptcha", service.Server(), new(subscriber.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
