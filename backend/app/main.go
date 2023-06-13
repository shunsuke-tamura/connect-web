package main

import (
	"context"
	"log"
	"net/http"

	"connect-back/db"
	taskv1 "connect-back/gen/rpc/task/v1"
	"connect-back/gen/rpc/task/v1/taskv1connect"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TaskListServer struct{}

func (s *TaskListServer) GetTaskList(ctx context.Context, req *connect_go.Request[taskv1.GetTaskListRequest]) (*connect_go.Response[taskv1.GetTaskListResponse], error) {
	log.Println("Request headers: ", req.Header())
	testTasks := []*taskv1.Task{
		{
			Id:          "1",
			UserId:      "123",
			Name:        "GetTaskList",
			IdCompleted: true,
		},
	}
	res := connect.NewResponse(&taskv1.GetTaskListResponse{
		Tasks: testTasks,
	})
	return res, nil
}

func (s *TaskListServer) CreateTask(ctx context.Context, req *connect_go.Request[taskv1.CreateTaskRequest]) (*connect_go.Response[taskv1.CreateTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&taskv1.CreateTaskResponse{
		CreatedId: "CreateTask",
	})
	return res, nil
}

func (s *TaskListServer) CompleteTask(ctx context.Context, req *connect_go.Request[taskv1.CompleteTaskRequest]) (*connect_go.Response[taskv1.CompleteTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&taskv1.CompleteTaskResponse{})
	return res, nil
}

func (s *TaskListServer) DeleteTask(ctx context.Context, req *connect_go.Request[taskv1.DeleteTaskRequest]) (*connect_go.Response[taskv1.DeleteTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&taskv1.DeleteTaskResponse{})
	return res, nil
}

func main() {
	db.Init()
	s := &TaskListServer{}
	mux := http.NewServeMux()
	path, handler := taskv1connect.NewTaskServiceHandler(s)
	mux.Handle(path, handler)
	log.Println("server is launched")
	http.ListenAndServe(
		"0.0.0.0:8080",
		cors.AllowAll().Handler(
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
	defer db.Db.Close()
}
