package proto

import (
	"context"
	gh "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct{}

func (h HealthServer) Check(ctx context.Context, request *gh.HealthCheckRequest) (*gh.HealthCheckResponse, error) {
	return &gh.HealthCheckResponse{
		Status: gh.HealthCheckResponse_SERVING,
	}, nil
}

func (h HealthServer) Watch(request *gh.HealthCheckRequest, server gh.Health_WatchServer) error {
	return nil
}