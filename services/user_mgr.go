package services

import (
	"fmt"
	"sync"
	"time"

	protocols "github.com/tap2joy/Protocols/go/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var gUserMgr *UserMgr
var gUserOnce sync.Once

// GetServiceMgr 获取单例
func GetUserMgr() *UserMgr {
	gUserOnce.Do(func() {
		if gUserMgr == nil {
			gUserMgr = NewUserMgr()
		}
	})
	return gUserMgr
}

type UserInfo struct {
	Name       string // 用户名
	Gate       string // 网关地址
	OnlineTime int64  // 上线时间
	Channel    uint32 // 当前所在频道
}

type UserMgr struct {
	Users map[string]*UserInfo
}

func NewUserMgr() *UserMgr {
	mgr := &UserMgr{
		Users: make(map[string]*UserInfo),
	}
	return mgr
}

// 玩家上线
func (mgr *UserMgr) UserOnline(name string, gate string, channelId uint32) (*UserInfo, error) {
	var oldUser *UserInfo
	if _, ok := mgr.Users[name]; ok {
		if mgr.Users[name].Gate == gate {
			oldUser = mgr.Users[name]
		}
	}

	// update
	curTime := time.Now().Unix()
	mgr.Users[name] = &UserInfo{Name: name, Gate: gate, OnlineTime: curTime, Channel: channelId}
	fmt.Printf("user %d:%s online on gate %s\n", channelId, name, gate)
	return oldUser, nil
}

// 玩家下线
func (mgr *UserMgr) UserOffline(name string) error {
	if _, ok := mgr.Users[name]; !ok {
		return status.Errorf(codes.Code(protocols.ErrorCode_USER_NOT_EXIST_ERROR), "user not exist")
	}

	delete(mgr.Users, name)
	fmt.Printf("user %s offline\n", name)
	return nil
}

// 获取在线玩家列表
func (mgr *UserMgr) GetOnlineUsers(channelId uint32) ([]*UserInfo, error) {
	retUsers := make([]*UserInfo, 0)

	for _, v := range mgr.Users {
		if v.Channel == channelId {
			retUsers = append(retUsers, v)
		}
	}

	return retUsers, nil
}

// 获取指定玩家的在线时长，单位秒
func (mgr *UserMgr) GetUserDuration(name string) (uint32, error) {
	if _, ok := mgr.Users[name]; ok {
		curTime := time.Now().Unix()
		duration := uint32(curTime - mgr.Users[name].OnlineTime)
		return duration, nil
	}

	return 0, status.Errorf(codes.Code(protocols.ErrorCode_USER_NOT_EXIST_ERROR), "user not exist")
}

// 修改用户所在频道
func (mgr *UserMgr) ChangeUserChannel(name string, channel uint32) error {
	if _, ok := mgr.Users[name]; ok {
		mgr.Users[name].Channel = channel
	}

	fmt.Printf("user %s switch to channel %d\n", name, channel)
	return nil
}
