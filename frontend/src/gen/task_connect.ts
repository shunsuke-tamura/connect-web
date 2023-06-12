// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=ts"
// @generated from file task.proto (package rpc.task.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CompleteTaskRequest, CompleteTaskResponse, CreateTaskRequest, CreateTaskResponse, DeleteTaskRequest, DeleteTaskResponse, GetTaskListRequest, GetTaskListResponse } from "./task_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service rpc.task.v1.TaskService
 */
export const TaskService = {
  typeName: "rpc.task.v1.TaskService",
  methods: {
    /**
     * @generated from rpc rpc.task.v1.TaskService.GetTaskList
     */
    getTaskList: {
      name: "GetTaskList",
      I: GetTaskListRequest,
      O: GetTaskListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc rpc.task.v1.TaskService.CreateTask
     */
    createTask: {
      name: "CreateTask",
      I: CreateTaskRequest,
      O: CreateTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc rpc.task.v1.TaskService.CompleteTask
     */
    completeTask: {
      name: "CompleteTask",
      I: CompleteTaskRequest,
      O: CompleteTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc rpc.task.v1.TaskService.DeleteTask
     */
    deleteTask: {
      name: "DeleteTask",
      I: DeleteTaskRequest,
      O: DeleteTaskResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
