package transportgrpc

import (
	"context"
	"fmt"
	taskpb "github.com/p1maf/grpcprot/proto/task"
	userpb "github.com/p1maf/grpcprot/proto/user"
	"github.com/p1maf/task-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, userClient userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: userClient}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.Userid}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.Userid, err)
	}

	t, err := h.svc.CreateTask(task.Task{
		UserId: req.Userid,
		Title:  req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.Id),
			Userid: t.UserId,
			Title:  t.Title,
		},
	}, nil
}

// Остальные CRUD методы

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.Id),
			Userid: t.UserId,
			Title:  t.Title,
		},
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, _ *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasks()
	if err != nil {
		return nil, err
	}
	var resp []*taskpb.Task
	for _, t := range tasks {
		resp = append(resp, &taskpb.Task{
			Id:     uint32(t.Id),
			Userid: t.UserId,
			Title:  t.Title,
		})
	}
	return &taskpb.ListTasksResponse{Tasks: resp}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.Userid}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.Userid, err)
	}

	t, err := h.svc.UpdateTask(task.Task{
		Id:     int(req.Id),
		UserId: req.Userid,
		Title:  req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &taskpb.UpdateTaskResponse{Task: &taskpb.Task{
		Id:     uint32(t.Id),
		Userid: t.UserId,
		Title:  t.Title,
	}}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(req.Id); err != nil {
		return nil, err
	}
	return &taskpb.DeleteTaskResponse{}, nil
}
