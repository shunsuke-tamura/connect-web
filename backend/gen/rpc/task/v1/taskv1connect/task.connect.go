// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc/task/v1/task.proto

package taskv1connect

import (
	v1 "connect-back/gen/rpc/task/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "rpc.task.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceGetTaskListProcedure is the fully-qualified name of the TaskService's GetTaskList RPC.
	TaskServiceGetTaskListProcedure = "/rpc.task.v1.TaskService/GetTaskList"
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/rpc.task.v1.TaskService/CreateTask"
	// TaskServiceCompleteTaskProcedure is the fully-qualified name of the TaskService's CompleteTask
	// RPC.
	TaskServiceCompleteTaskProcedure = "/rpc.task.v1.TaskService/CompleteTask"
	// TaskServiceDeleteTaskProcedure is the fully-qualified name of the TaskService's DeleteTask RPC.
	TaskServiceDeleteTaskProcedure = "/rpc.task.v1.TaskService/DeleteTask"
)

// TaskServiceClient is a client for the rpc.task.v1.TaskService service.
type TaskServiceClient interface {
	GetTaskList(context.Context, *connect_go.Request[v1.GetTaskListRequest]) (*connect_go.Response[v1.GetTaskListResponse], error)
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error)
}

// NewTaskServiceClient constructs a client for the rpc.task.v1.TaskService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		getTaskList: connect_go.NewClient[v1.GetTaskListRequest, v1.GetTaskListResponse](
			httpClient,
			baseURL+TaskServiceGetTaskListProcedure,
			opts...,
		),
		createTask: connect_go.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			opts...,
		),
		completeTask: connect_go.NewClient[v1.CompleteTaskRequest, v1.CompleteTaskResponse](
			httpClient,
			baseURL+TaskServiceCompleteTaskProcedure,
			opts...,
		),
		deleteTask: connect_go.NewClient[v1.DeleteTaskRequest, v1.DeleteTaskResponse](
			httpClient,
			baseURL+TaskServiceDeleteTaskProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	getTaskList  *connect_go.Client[v1.GetTaskListRequest, v1.GetTaskListResponse]
	createTask   *connect_go.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	completeTask *connect_go.Client[v1.CompleteTaskRequest, v1.CompleteTaskResponse]
	deleteTask   *connect_go.Client[v1.DeleteTaskRequest, v1.DeleteTaskResponse]
}

// GetTaskList calls rpc.task.v1.TaskService.GetTaskList.
func (c *taskServiceClient) GetTaskList(ctx context.Context, req *connect_go.Request[v1.GetTaskListRequest]) (*connect_go.Response[v1.GetTaskListResponse], error) {
	return c.getTaskList.CallUnary(ctx, req)
}

// CreateTask calls rpc.task.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// CompleteTask calls rpc.task.v1.TaskService.CompleteTask.
func (c *taskServiceClient) CompleteTask(ctx context.Context, req *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error) {
	return c.completeTask.CallUnary(ctx, req)
}

// DeleteTask calls rpc.task.v1.TaskService.DeleteTask.
func (c *taskServiceClient) DeleteTask(ctx context.Context, req *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the rpc.task.v1.TaskService service.
type TaskServiceHandler interface {
	GetTaskList(context.Context, *connect_go.Request[v1.GetTaskListRequest]) (*connect_go.Response[v1.GetTaskListResponse], error)
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(TaskServiceGetTaskListProcedure, connect_go.NewUnaryHandler(
		TaskServiceGetTaskListProcedure,
		svc.GetTaskList,
		opts...,
	))
	mux.Handle(TaskServiceCreateTaskProcedure, connect_go.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	))
	mux.Handle(TaskServiceCompleteTaskProcedure, connect_go.NewUnaryHandler(
		TaskServiceCompleteTaskProcedure,
		svc.CompleteTask,
		opts...,
	))
	mux.Handle(TaskServiceDeleteTaskProcedure, connect_go.NewUnaryHandler(
		TaskServiceDeleteTaskProcedure,
		svc.DeleteTask,
		opts...,
	))
	return "/rpc.task.v1.TaskService/", mux
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) GetTaskList(context.Context, *connect_go.Request[v1.GetTaskListRequest]) (*connect_go.Response[v1.GetTaskListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rpc.task.v1.TaskService.GetTaskList is not implemented"))
}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rpc.task.v1.TaskService.CreateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rpc.task.v1.TaskService.CompleteTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rpc.task.v1.TaskService.DeleteTask is not implemented"))
}
