package server

import (
	"context"
	"time"

	"github.com/tap2joy/CenterService/services"
	pb "github.com/tap2joy/Protocols/go/grpc/center"
)

type Server struct {
}

// 用户上线
func (*Server) UserOnline(ctx context.Context, req *pb.UserOnlineRequest) (*pb.UserOnlineResponse, error) {
	name := req.Name
	gate := req.Gate
	channel := req.Channel

	oldUser, err := services.GetUserMgr().UserOnline(name, gate, channel)
	if err != nil {
		return nil, err
	}

	resp := new(pb.UserOnlineResponse)
	if oldUser != nil {
		resp.OldUser = &pb.UserInfo{Name: oldUser.Name, Gateway: oldUser.Gate}
	}

	return resp, nil
}

// 用户下线
func (*Server) UserOffline(ctx context.Context, req *pb.UserOfflineRequest) (*pb.UserOfflineResponse, error) {
	name := req.Name

	err := services.GetUserMgr().UserOffline(name)
	if err != nil {
		return nil, err
	}

	resp := new(pb.UserOfflineResponse)
	return resp, nil
}

// 注册服务
func (*Server) RegisterService(ctx context.Context, req *pb.RegisterServiceRequest) (*pb.RegisterServiceResponse, error) {
	serverType := req.Type
	address := req.Address

	err := services.GetServiceMgr().RegisterService(serverType, address)
	if err != nil {
		return nil, err
	}

	resp := new(pb.RegisterServiceResponse)
	return resp, nil
}

// 获取可用服务列表
func (*Server) GetServices(ctx context.Context, req *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	serverType := req.Type

	list, err := services.GetServiceMgr().GetServices(serverType)
	if err != nil {
		return nil, err
	}

	resp := new(pb.GetServicesResponse)
	for _, v := range list {
		resp.List = append(resp.List, v.Address)
	}

	return resp, nil
}

// 获取在线用户列表
func (*Server) GetOnlineUsers(ctx context.Context, req *pb.GetOnlineUsersRequest) (*pb.GetOnlineUsersResponse, error) {
	channelId := req.Channel

	userList, err := services.GetUserMgr().GetOnlineUsers(channelId)
	if err != nil {
		return nil, err
	}

	resp := new(pb.GetOnlineUsersResponse)
	for _, v := range userList {
		resp.Users = append(resp.Users, &pb.UserInfo{Name: v.Name, Gateway: v.Gate})
	}

	return resp, nil
}

// 获取用户本次在线时长
func (*Server) GetUserOnlineTime(ctx context.Context, req *pb.GetUserOnlineTimeRequest) (*pb.GetUserOnlineTimeResponse, error) {
	name := req.Name

	duration, err := services.GetUserMgr().GetUserDuration(name)
	if err != nil {
		return nil, err
	}

	resp := new(pb.GetUserOnlineTimeResponse)
	resp.Duration = duration
	return resp, nil
}

// 心跳
func (*Server) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	serverType := req.Type
	address := req.Address

	services.GetServiceMgr().HeartBeat(serverType, address)

	resp := new(pb.HeartBeatResponse)
	resp.Timestamp = time.Now().Unix()
	return resp, nil
}

// 切换频道
func (*Server) ChangeChannel(ctx context.Context, req *pb.ChangeChannelRequest) (*pb.ChangeChannelResponse, error) {
	name := req.Name
	channel := req.Channel

	services.GetUserMgr().ChangeUserChannel(name, channel)

	resp := new(pb.ChangeChannelResponse)
	return resp, nil
}
