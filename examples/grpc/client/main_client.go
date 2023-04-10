package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lark/pkg/proto/pb_auth"
)

func main() {
	req := &pb_auth.SignUpReq{
		Nickname: "张三",
	}
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb_auth.NewAuthClient(conn)
	var resp *pb_auth.SignUpResp
	resp, err = client.SignUp(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if resp == nil {
		return
	}
	fmt.Println(resp)
}
