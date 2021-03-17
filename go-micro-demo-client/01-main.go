package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	proto "go-micro-demo-client/proto/go-micro-demo"
	"golang.org/x/net/context"
)

func main() {
	etcd := etcd.NewRegistry(
		// 地址是我本地etcd服务器地址，不要照抄
		registry.Addrs("localhost:2379"),
	)
	service := micro.NewService(
		micro.Name("go.micro.client"),
		micro.Version("latest"),
		micro.Registry(etcd),
	)

	service.Init()

	client := proto.NewGetCaptchaService("go.micro.service.getCaptcha", service.Client())

	rsp, err := client.Call(context.TODO(), &proto.Request{
		Uuid: "123456",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Img)
}
