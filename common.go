package proto

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/sercand/kuberesolver/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var entry = logrus.NewEntry(logrus.StandardLogger())

func GetLogger() *logrus.Entry {
	return entry
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
