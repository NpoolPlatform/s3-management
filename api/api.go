package api

import (
	"context"

	"github.com/NpoolPlatform/s3-management/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedS3ManagementServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterS3ManagementServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterS3ManagementHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
