package client

import (
	"context"
	"fmt"
	pb "github.com/Abubakarr99/taskManager/proto"
	"github.com/Abubakarr99/taskManager/storage/boltdb"
	"google.golang.org/grpc"
	"io"
)

// Client is a client to the task service

type Client struct {
	client pb.TaskManagerClient
	conn   *grpc.ClientConn
}

// New is the constructor for client
func New(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		client: pb.NewTaskManagerClient(conn),
		conn:   conn,
	}, nil
}

// CallOptions options for the RPC call.
type CallOptions func(co *callOptions)

type callOptions struct {
	trace *string
}

// Task is wrapper around
type Task struct {
	*pb.Task
	err error
}

func (c *Client) AddTasks(ctx context.Context, tasks []*pb.Task) ([]string, error) {
	// check is tasks is ont empty
	if len(tasks) == 0 {
		return nil, nil
	}
	for _, task := range tasks {
		if err := boltdb.Validate(task); err != nil {
			return nil, err
		}
	}
	resp, err := c.client.AddTasks(ctx, &pb.AddTaskReq{Tasks: tasks})
	if err != nil {
		return nil, err
	}
	return resp.Ids, nil
}

func (c *Client) UpdateTasks(ctx context.Context, tasks []*pb.Task) error {
	// check is tasks is not empty
	if len(tasks) == 0 {
		return nil
	}
	for _, task := range tasks {
		if err := boltdb.Validate(task); err != nil {
			return err
		}
	}
	_, err := c.client.UpdateTasks(ctx, &pb.UpdateTasksReq{Tasks: tasks})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteTasks(ctx context.Context, ids []string) error {
	// check is tasks is not empty
	if len(ids) == 0 {
		return nil
	}
	_, err := c.client.DeleteTasks(ctx, &pb.DeleteTasksReq{Ids: ids})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SearchTasks(ctx context.Context, filter *pb.SearchTaskReq) (chan Task, error) {
	if filter == nil {
		return nil, fmt.Errorf("the filter cannot be empty")
	}
	stream, err := c.client.SearchTasks(ctx, filter)
	if err != nil {
		return nil, err
	}
	ch := make(chan Task, 1)
	go func() {
		defer close(ch)
		for {
			t, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				ch <- Task{err: err}
			}
			ch <- Task{Task: t}
		}
	}()
	return ch, nil
}

func (t Task) Error() error {
	return t.err
}
