package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/CenterService/services"
)

func TestServiceMgr(t *testing.T) {
	serviceMgr := services.NewServiceMgr()
	serviceMgr.RegisterService("chat", "127.0.0.1:9101")
	serviceMgr.RegisterService("chat", "127.0.0.1:9102")

	services, err := serviceMgr.GetServices("chat")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, v := range services {
		fmt.Printf("%v\n", v)
	}
}
