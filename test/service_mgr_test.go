package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tap2joy/CenterService/services"
)

func TestServiceMgr(t *testing.T) {
	fmt.Println("======================================= ServiceMgr test begin")

	serviceMgr := services.NewServiceMgr()
	serviceMgr.RegisterService("chat", "127.0.0.1:9101")
	serviceMgr.RegisterService("chat", "127.0.0.1:9102")
	serviceMgr.RegisterService("chat", "127.0.0.1:9103")

	fmt.Println("current chat services:")
	services, err := serviceMgr.GetServices("chat")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, v := range services {
		fmt.Printf("%v\n", v)
	}

	time.Sleep(time.Duration(6) * time.Second)
	serviceMgr.HeartBeat("chat", "127.0.0.1:9103")

	fmt.Println("After clean timeout service, current chat services:")
	services, err = serviceMgr.GetServices("chat")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, v := range services {
		fmt.Printf("%v\n", v)
	}
}
