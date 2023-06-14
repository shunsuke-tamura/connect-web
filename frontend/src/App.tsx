import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

import { useClient } from "./common/use-client";
import { TaskService } from "./gen/task_connect";
import { Task } from "./gen/task_pb";
import { JsonValue } from "@bufbuild/protobuf";

function App() {
  const [count, setCount] = useState(0);
  const [tasks, setTasks] = useState<JsonValue>("");

  useEffect(() => {
    console.log("effect", tasks);
  });

  const c = useClient(TaskService);
  const send = async () => {
    const stream = c.getTaskList({});
    for await (const res of stream) {
      res.toJson() === tasks ?? setTasks(res.toJson());
      console.log("stream", res.toJson());
    }
  };
  send();
  // useClient(TaskService)
  //   .getTaskList({})
  //   .then((res) => {
  //     console.log(res.toJson());
  //   });
  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
