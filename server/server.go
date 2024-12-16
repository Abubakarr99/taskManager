// Package server contains the grpc implementation of the taskManager server
package server

import (
	"context"
	"fmt"
	pb "github.com/Abubakarr99/taskManager/proto"
	"github.com/Abubakarr99/taskManager/storage/boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	return &pb.UpdateTasksResp{}, nil
}

func (a *API) DeleteTasks(ctx context.Context, req *pb.DeleteTasksReq) (*pb.DeleteTaskResp, error) {
	if err := a.db.DeleteTasks(ctx, req.Ids); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteTaskResp{}, nil
}
