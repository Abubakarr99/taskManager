package boltdb

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/Abubakarr99/taskManager/proto"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
	"strings"
	"sync"
	"time"
)

var (
	taskBucket = []byte("tasks")
	//db         *bolt.DB
)

type TaskManagerServer struct {
	Db *bolt.DB
	pb.UnimplementedTaskManagerServer
	mu      sync.RWMutex
	ids     map[string]*pb.Task
	titles  map[string]map[string]*pb.Task
	urgency map[pb.Urgency]map[string]*pb.Task
}

func Init(dbPath string) (*TaskManagerServer, error) {
	database, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = database.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("failed to create tasks bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &TaskManagerServer{
		Db:      database,
		ids:     make(map[string]*pb.Task),
		titles:  make(map[string]map[string]*pb.Task),
		urgency: make(map[pb.Urgency]map[string]*pb.Task),
	}, nil
}

func Validate(task *pb.Task) error {
	if task.Id == "" {
		return fmt.Errorf("cannot have a task without ID")
	}
	title := strings.TrimSpace(task.Title)
	if title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if _, valid := pb.Urgency_name[int32(task.Urgency)]; !valid {
		return fmt.Errorf("invalid urgency level: %v", task.Urgency)
	}
	return nil
}

func (s *TaskManagerServer) Validate(task *pb.Task) error {
	if task.Id == "" {
		return fmt.Errorf("cannot have a task without ID")
	}
	title := strings.TrimSpace(task.Title)
	if title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if _, valid := pb.Urgency_name[int32(task.Urgency)]; !valid {
		return fmt.Errorf("invalid urgency level: %v", task.Urgency)
	}
	return nil
}

// AddTasks add tasks to the bolt DB
func (s *TaskManagerServer) AddTasks(tasks []*pb.Task) ([]string, error) {
	var ids []string
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		for _, task := range tasks {
			task.Id = uuid.New().String()
			taskjson, err := json.Marshal(task)
			if err != nil {
				return fmt.Errorf("failed to serialize task: %w", err)
			}
			if err := b.Put([]byte(task.Id), taskjson); err != nil {
				return fmt.Errorf("failed to add task to db: %w", err)
			}
			ids = append(ids, task.Id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// UpdateTasks updates tasks in the BD
func (s *TaskManagerServer) UpdateTasks(tasks []*pb.Task) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	err := s.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		for _, task := range tasks {
			taskJson, err := json.Marshal(task)
			if err != nil {
				return fmt.Errorf("failed to serialise task: %w", err)
			}
			if err := b.Put([]byte(task.Id), taskJson); err != nil {
				return fmt.Errorf("failed to update task: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteTasks deletes tasks from the database
func (s *TaskManagerServer) DeleteTasks(ctx context.Context, ids []string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	err := s.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		for _, id := range ids {
			if err := b.Delete([]byte(id)); err != nil {
				return fmt.Errorf("failed to delete task %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskManagerServer) SearchTasks(req *pb.SearchTaskReq) ([]*pb.Task, error) {
	var matchedTasks []*pb.Task
	err := s.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.ForEach(func(k, v []byte) error {
			var task pb.Task
			if err := json.Unmarshal(v, &task); err != nil {
				return fmt.Errorf("failed to deserialize task: %w", err)
			}

			// Filter by urgency of the task
			if len(req.Urgency) > 0 {
				found := false
				for _, u := range req.Urgency {
					if task.Urgency == u {
						found = true
						break
					}
				}
				if !found {
					return nil
				}
			}
			// Filter by date range
			if req.DueRange != nil {
				if !isDateInRange(task.DueDate, req.DueRange) {
					return nil
				}
			}
			matchedTasks = append(matchedTasks, &task)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return matchedTasks, nil
}

// Helper to check if a date is in range
func isDateInRange(dueDate *date.Date, rangeFilter *pb.DateRange) bool {
	due := time.Date(int(dueDate.Year), time.Month(dueDate.Month), int(dueDate.Day), 0, 0, 0, 0, time.UTC)
	start := time.Date(int(rangeFilter.Start.Year), time.Month(rangeFilter.Start.Month), int(rangeFilter.Start.Day), 0, 0, 0, 0, time.UTC)
	end := time.Date(int(rangeFilter.End.Year), time.Month(rangeFilter.End.Month), int(rangeFilter.End.Day), 0, 0, 0, 0, time.UTC)
	return due.After(start) && due.Before(end)
}
