package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"lark/examples/pb_auth"
	"net"
)

func main() {
	authServer := new(AuthServer)
	authServer.Run()
}

type AuthServer struct {
	pb_auth.UnimplementedAuthServer
}

func (s *AuthServer) Run() {
	var (
		listener net.Listener
		server   *grpc.Server
		err      error
	)
	// 传入一个监听
	listener, err = net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	server = grpc.NewServer()
	pb_auth.RegisterAuthServer(server, s)
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	fmt.Println(req.NickName)
	resp = new(pb_auth.SignUpResp)
	resp.Msg = "注册成功"
	return
}
