package xgrpc

import (
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"lark/pkg/common/xetcd"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
	"net"
	"strconv"
)

type GrpcServer struct {
	grpcConf *conf.Grpc
	etcdConf *conf.Etcd
}

func NewGrpcServer(grpc *conf.Grpc, etcd *conf.Etcd) *GrpcServer {
	return &GrpcServer{grpc, etcd}
}

func (s *GrpcServer) RunServer(server *grpc.Server) {
	var (
		address  string
		listener net.Listener
		err      error
	)
	defer func() {
		server.GracefulStop()
	}()

	address = "0.0.0.0:" + strconv.Itoa(s.grpcConf.Port)
	listener, err = net.Listen("tcp", address)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	if s.grpcConf.ConnectionLimit > 0 {
		// 最大并发连接数
		listener = netutil.LimitListener(listener, s.grpcConf.ConnectionLimit)
	}
	err = xetcd.RegisterEtcd(s.etcdConf.Schema, s.etcdConf.Endpoints, utils.GetServerIP(), s.grpcConf.Port, s.grpcConf.Name, 10)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	err = server.Serve(listener)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
}
