package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connect-back/db"
	taskv1 "connect-back/gen/rpc/task/v1"
	"connect-back/gen/rpc/task/v1/taskv1connect"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var stream *connect_go.ServerStream[taskv1.GetTaskListResponse]

type TaskListServer struct{}

func fetchTasks() []*taskv1.Task {
	rows, err := db.Db.Query("select * from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tasks []*taskv1.Task
	for rows.Next() {
		var task taskv1.Task
		var createdAt time.Time
		var updatedAt time.Time
		err := rows.Scan(&task.Id, &task.UserId, &task.Name, &task.IsCompleted, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		task.CreatedAt = timestamppb.New(createdAt)
		task.UpdatedAt = timestamppb.New(updatedAt)
		tasks = append(tasks, &task)
	}
	return tasks
}

func (s *TaskListServer) GetTaskList(ctx context.Context, req *connect_go.Request[taskv1.GetTaskListRequest], stm *connect_go.ServerStream[taskv1.GetTaskListResponse]) error {
	log.Println("Request headers: ", req.Header())
	stream = stm
	stream.Send(&taskv1.GetTaskListResponse{
		Tasks: fetchTasks(),
	})
	time.Sleep(60 * time.Second)
	return nil
}

func (s *TaskListServer) CreateTask(ctx context.Context, req *connect_go.Request[taskv1.CreateTaskRequest]) (*connect_go.Response[taskv1.CreateTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	resu, err := db.Db.Exec("insert into tasks(user_id, name, is_completed) values ('123', ?, false)", req.Msg.Name)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := resu.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	res := connect.NewResponse(&taskv1.CreateTaskResponse{
		CreatedId: lastId,
	})
	stream.Send(&taskv1.GetTaskListResponse{
		Tasks: fetchTasks(),
	})
	return res, nil
}

func (s *TaskListServer) CompleteTask(ctx context.Context, req *connect_go.Request[taskv1.CompleteTaskRequest]) (*connect_go.Response[taskv1.CompleteTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	_, err := db.Db.Exec("update tasks set is_completed = true where id = ?", req.Msg.TaskId)
	if err != nil {
		log.Fatal(err)
	}
	res := connect.NewResponse(&taskv1.CompleteTaskResponse{})
	stream.Send(&taskv1.GetTaskListResponse{
		Tasks: fetchTasks(),
	})
	return res, nil
}

func (s *TaskListServer) DeleteTask(ctx context.Context, req *connect_go.Request[taskv1.DeleteTaskRequest]) (*connect_go.Response[taskv1.DeleteTaskResponse], error) {
	log.Println("Request headers: ", req.Header())
	_, err := db.Db.Exec("delete from tasks where id = ?", req.Msg.TaskId)
	if err != nil {
		log.Fatal(err)
	}
	res := connect.NewResponse(&taskv1.DeleteTaskResponse{})
	stream.Send(&taskv1.GetTaskListResponse{
		Tasks: fetchTasks(),
	})
	return res, nil
}

func main() {
	db.Init()
	s := &TaskListServer{}
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"rpc.task.v1.TaskService", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
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
