import { useClient } from "../common/use-client";
import { TaskService } from "../gen/task_connect";
import { CompleteTaskRequest, DeleteTaskRequest, Task } from "../gen/task_pb";

type Props = {
  task: Task;
};

const TaskCard = (props: Props) => {
  const c = useClient(TaskService);

  const containerStyle = {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    margin: "10px 0",
    padding: "10px",
    backgroundColor: "#f2f2f2",
    borderRadius: "5px",
    boxShadow: "2px 2px 5px rgba(0, 0, 0, 0.3)",
  };

  const onDelete = (id: bigint) => {
    console.log("delete");

    c.deleteTask(new DeleteTaskRequest({ taskId: id }));
  };
  const deleteStyle = {
    color: "#dc3545",
    cursor: "pointer",
  };

  const onComplete = (id: bigint) => {
    c.completeTask(new CompleteTaskRequest({ taskId: id }));
  };
  const completeStyle = {
    color: "#0BC46F",
    cursor: "pointer",
  };
  const completedStyle = {
    color: "#7414F8",
    cursor: "pointer",
  };

  return (
    <div className="todo" css={containerStyle}>
      <span style={{ color: "black" }}>{props.task.name}</span>
      {props.task.isCompleted ? (
        <span className="complete" css={completedStyle}>
          完了済み
        </span>
      ) : (
        <span
          className="complete"
          onClick={() => onComplete(props.task.id)}
          css={completeStyle}
        >
          to完了
        </span>
      )}
      <span
        className="delete"
        onClick={() => onDelete(props.task.id)}
        css={deleteStyle}
      >
        削除
      </span>
    </div>
  );
};

export default TaskCard;
