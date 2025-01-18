package server

import (
	"context"
	pb "github.com/Abubakarr99/taskManager/proto"
)

// HealthService implements the health gRPC interface
type HealthService struct {
	pb.UnimplementedHealthServer
}

func (h *HealthService) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

// Watch handles server-side streaming health checks
func (h *HealthService) Watch(req *pb.HealthCheckRequest, stream pb.Health_WatchServer) error {
	// Example: Stream SERVING status every second
	stream.Send(&pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	})
	return nil
}
