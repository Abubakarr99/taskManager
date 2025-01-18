// Package server contains the grpc implementation of the taskManager server
package server

import (
	"context"
	"fmt"
	pb "github.com/Abubakarr99/taskManager/proto"
	"github.com/Abubakarr99/taskManager/storage/boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

// API implements the gRPC server's API
type API struct {
	pb.UnimplementedTaskManagerServer
	db         *boltdb.TaskManagerServer
	addr       string
	grpcServer *grpc.Server
	gOpts      []grpc.ServerOption
	mu         sync.Mutex
}

// Option is an optional arguments to New()
type Option func(a *API)

// WithGRPCOpts creates the gRPC server with the options passed
func WithGRPCOpts(opts ...grpc.ServerOption) Option {
	return func(a *API) {
		a.gOpts = append(a.gOpts, opts...)
	}
}

// New is the constructor for the API
func New(addr string, db *boltdb.TaskManagerServer, options ...Option) (*API, error) {
	a := &API{addr: addr, db: db}
	for _, o := range options {
		o(a)
	}
	a.grpcServer = grpc.NewServer(a.gOpts...)
	a.grpcServer.RegisterService(&pb.TaskManager_ServiceDesc, a)
	// register the healthServer
	healthServer := health.NewServer()
	healthServer.SetServingStatus("TaskManager", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(a.grpcServer, healthServer)
	reflection.Register(a.grpcServer)
	return a, nil
}

// Start the server
func (a *API) Start() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	listen, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}
	fmt.Println("gRPC server listening on", a.addr)
	return a.grpcServer.Serve(listen)
}

// Stop the server
func (a *API) Stop() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.grpcServer.Stop()
}

// AddTasks adds task manager
func (a *API) AddTasks(ctx context.Context, req *pb.AddTaskReq) (*pb.AddTasksResp, error) {
	for _, task := range req.Tasks {
		if err := a.db.Validate(task); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	ids, err := a.db.AddTasks(req.Tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	log.Printf("successfully added the tasks %s", ids)
	return &pb.AddTasksResp{Ids: ids}, nil
}

// UpdateTasks add tasks to the task manager
func (a *API) UpdateTasks(ctx context.Context, req *pb.UpdateTasksReq) (*pb.UpdateTasksResp, error) {
	for _, task := range req.Tasks {
		if err := a.db.Validate(task); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	err := a.db.UpdateTasks(req.Tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	log.Printf("successfully updated the task %s", req.Tasks[0].Id)
	return &pb.UpdateTasksResp{}, nil
}

func (a *API) DeleteTasks(ctx context.Context, req *pb.DeleteTasksReq) (*pb.DeleteTaskResp, error) {
	if err := a.db.DeleteTasks(ctx, req.Ids); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	log.Printf("successfully deleted the tasks %s", req.Ids)
	return &pb.DeleteTaskResp{}, nil
}

func (a *API) SearchTasks(req *pb.SearchTaskReq, stream pb.TaskManager_SearchTasksServer) error {
	// Call the DB-side method to retrieve matching tasks
	tasks, err := a.db.SearchTasks(req)
	if err != nil {
		return fmt.Errorf("failed to search tasks: %w", err)
	}

	// Stream tasks back to the client
	for _, task := range tasks {
		if err := stream.Send(task); err != nil {
			return fmt.Errorf("failed to stream task: %w", err)
		}
	}
	return nil
}
