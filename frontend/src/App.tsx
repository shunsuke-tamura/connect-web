import { Suspense, useEffect, useState } from "react";

import "./App.css";

import { useClient } from "./common/use-client";
import { TaskService } from "./gen/task_connect";
import { CreateTaskRequest, Task } from "./gen/task_pb";
import TaskCard from "./components/TaskCard";

function App() {
  const [title, setTitle] = useState("");
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    const now = new Date();
    console.log(
      `${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}`,
      tasks
    );
  });

  const c = useClient(TaskService);
  const send = async () => {
    const stream = c.getTaskList({});
    for await (const res of stream) {
      if (res.tasks !== tasks) setTasks(res.tasks);
    }
  };

  useEffect(() => {
    // send();
  });

  const addTask = () => {
    c.createTask(new CreateTaskRequest({ name: title }));
  };

  return (
    <>
      <h1>ToDoアプリ</h1>
      <button onClick={() => send()}>fetch</button>
      <form
        onSubmit={undefined}
        onKeyDown={(e) => e.key === "Enter" ?? addTask()}
      >
        <input
          type="text"
          placeholder="タスクを入力"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <button type="button" onClick={() => addTask()}>
          追加
        </button>
      </form>
      <Suspense fallback={<p>Loading...</p>}>
        <div>
          {tasks.map((task) => {
            return <TaskCard task={task} key={Number(task.id)}></TaskCard>;
          })}
        </div>
      </Suspense>
    </>
  );
}

export default App;
