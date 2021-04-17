package services

import (
	"fmt"
	"sync"
	"time"
)

var instance *ServiceMgr
var once sync.Once

// GetServiceMgr 获取单例
func GetServiceMgr() *ServiceMgr {
	once.Do(func() {
		if instance == nil {
			instance = NewServiceMgr()
		}
	})
	return instance
}

type ServiceInfo struct {
	Type          string // 服务类型
	Address       string // 服务地址
	HeartBeatTime int64  // 最近一次心跳时间
}

type ServiceMgr struct {
	Services      []*ServiceInfo // 服务列表
	LastCleanTime int64          // 上一次清理超时服务的时间
}

func NewServiceMgr() *ServiceMgr {
	mgr := &ServiceMgr{
		LastCleanTime: time.Now().Unix(),
	}
	return mgr
}

// 注册服务
func (mgr *ServiceMgr) RegisterService(serviceType string, address string) error {
	for _, info := range mgr.Services {
		if info.Type == serviceType && info.Address == address {
			// already exist, do nothing
			return nil //status.Errorf(codes.Code(protocols.ErrorCode_SERVICE_EXIST_ERROR), "service exist")
		}
	}

	mgr.Services = append(mgr.Services, &ServiceInfo{Type: serviceType, Address: address, HeartBeatTime: time.Now().Unix()})
	fmt.Printf("type: %s, address: %s register success\n", serviceType, address)
	return nil
}

// 获取指定类型的服务列表
func (mgr *ServiceMgr) GetServices(serviceType string) ([]*ServiceInfo, error) {
	var retServices []*ServiceInfo
	for _, info := range mgr.Services {
		if info.Type == serviceType {
			retServices = append(retServices, info)
		}
	}
	return retServices, nil
}

// 心跳同步
func (mgr *ServiceMgr) HeartBeat(serviceType string, address string) {
	exist := false
	for _, info := range mgr.Services {
		if info.Type == serviceType && info.Address == address {
			info.HeartBeatTime = time.Now().Unix()
			exist = true
			break
		}
	}

	if !exist {
		mgr.RegisterService(serviceType, address)
	}
	//fmt.Printf("heart beat, type: %s, address: %s\n", serviceType, address)
	mgr.CleanTimeoutService()
}

// 清理心跳超时的服务
func (mgr *ServiceMgr) CleanTimeoutService() {
	curTime := time.Now().Unix()
	if mgr.LastCleanTime+5 > curTime {
		return
	}

	for i := 0; i < len(mgr.Services); {
		if mgr.Services[i].HeartBeatTime+3 < curTime {
			fmt.Printf("remove timeout service [%s:%s]\n", mgr.Services[i].Type, mgr.Services[i].Address)
			mgr.Services = append(mgr.Services[:i], mgr.Services[i+1:]...)
		} else {
			i++
		}
	}

	mgr.LastCleanTime = time.Now().Unix()
}
