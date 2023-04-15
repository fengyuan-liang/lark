package main

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lark/examples/auth/constant"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
)

func main() {
	var (
		target = constant.GRPC_SERVER_ADDR
		conn   *grpc.ClientConn
		err    error
		client pb_auth.AuthClient
		req    *pb_auth.SignUpReq
		resp   *pb_auth.SignUpResp
	)
	req = &pb_auth.SignUpReq{
		RegPlatform: pb_enum.PLATFORM_TYPE_WINDOWS,
		Nickname:    "张三",
		Password:    "123456",
		Firstname:   "张",
		Lastname:    "三",
		Gender:      constant.MAN,
		Email:       "zhangsan@gmail.com",
		Code:        1234,
		Udid:        uuid.NewV4().String(),
		ServerId:    1,
	}
	// 获取grpc连接
	conn, err = grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client = pb_auth.NewAuthClient(conn)
	resp, err = client.SignUp(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if resp == nil {
		xlog.Error("服务器故障")
		return
	}
	xlog.Infof("resp[%v]", resp.Msg)
}
