package transportgrpc

import (
	taskpb "github.com/p1maf/grpcprot/proto/task"
	userpb "github.com/p1maf/grpcprot/proto/user"
	"github.com/p1maf/task-service/internal/task"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()

	taskpb.RegisterTaskServiceServer(grpcSrv, NewHandler(svc, uc))
	return grpcSrv.Serve(lis)
}
