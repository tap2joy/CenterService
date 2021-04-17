package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/CenterService/services"
)

func TestUserMgr(t *testing.T) {
	fmt.Println("======================================= UserMgr test begin")
	userMgr := services.NewUserMgr()
	userMgr.UserOnline("tom", "127.0.0.1:9101", 1)
	userMgr.UserOnline("lucy", "127.0.0.1:9101", 1)
	userMgr.UserOnline("jack", "127.0.0.1:9102", 1)

	onlineUsers, err := userMgr.GetOnlineUsers(1)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, v := range onlineUsers {
			fmt.Printf("%v\n", v)
		}
	}

	userMgr.UserOffline("tom")
	fmt.Println("after tom offline")
	onlineUsers, err = userMgr.GetOnlineUsers(1)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, v := range onlineUsers {
			fmt.Printf("%v\n", v)
		}
	}

	duration, err := userMgr.GetUserDuration("lucy")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("lucy online time is %d ms\n", duration)
	}
}
