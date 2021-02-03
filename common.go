package proto

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/sercand/kuberesolver/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"net"
)

var entry = logrus.NewEntry(logrus.StandardLogger())

func GetLogger() *logrus.Entry {
	return entry
}
func DefaultGrpcServer(serviceName string, chainStreamServer []grpc.StreamServerInterceptor, chainUnaryServer []grpc.UnaryServerInterceptor) *grpc.Server {
	grpc_logrus.ReplaceGrpcLogger(GetLogger())
	chainStreamServer = append(chainStreamServer, []grpc.StreamServerInterceptor{grpc_logrus.StreamServerInterceptor(GetLogger()),
		grpc_opentracing.StreamServerInterceptor(),
	}...)

	chainUnaryServer = append(chainUnaryServer, []grpc.UnaryServerInterceptor{grpc_logrus.UnaryServerInterceptor(GetLogger()),
		grpc_opentracing.UnaryServerInterceptor(),
	}...)
	return grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(chainStreamServer...)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(chainUnaryServer...)),
	)
}

func Run(s *grpc.Server, serverAddress string) error {
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}
	return s.Serve(lis)
}

func DefaultConn(target string) (*grpc.ClientConn, error) {
	kuberesolver.RegisterInCluster()
	return grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(),
			grpc_logrus.UnaryClientInterceptor(GetLogger()),
		),
		grpc.WithChainStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(),
			grpc_logrus.StreamClientInterceptor(GetLogger())),

		grpc.WithBalancerName(roundrobin.Name),
	)
}
