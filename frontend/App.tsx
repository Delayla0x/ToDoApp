import React, { useState } from 'react';

interface Task {
  id: number;
  text: string;
}

const TaskForm: React.FC<{ onAddTask: (taskText: string) => void }> = (props) => {
  const [enteredText, setEnteredText] = useState("");

  const addTaskHandler = (event: React.FormHTMLAttributes<HTMLFormElement>) => {
    event.preventDefault();
    if (enteredText.trim().length === 0) {
      return;
    }
    props.onAddTask(enteredText);
    setEnteredText("");
  };

  return (
    <form onSubmit={addTaskHandler}>
      <input
        type="text"
        value={enteredText}
        onChange={(event) => setEnteredText(event.target.value)}
      />
      <button type="submit">Add Task</button>
    </form>
  );
};

const TaskList: React.FC<{ items: Task[]; onDeleteTask: (taskId: number) => void; onEditTask: (taskId: number, newText: string) => void }> = (props) => {
  const [editState, setEditState] = useState<{ id: number | null, text: string }>({ id: null, text: "" });

  const startEditing = (task: Task) => {
    setEditState({ id: task.id, text: task.text });
  };

  const editTaskHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEditState((prevState) => ({...prevState, text: event.target.value}));
  };

  const submitEditTaskHandler = (taskId: number) => {
    props.onEditTask(taskId, editState.text);
    setEditState({ id: null, text: "" });
  };

  return (
    <ul>
      {props.items.map(task => (
        <li key={task.id}>
          {editState.id === task.id ? (
            <input
              type="text"
              value={editStat.text}
              onChange={editTaskHandler}
            />) : (
              task.text
            )}
          <button onClick={() => startEditing(task)}>Edit</button>
          <button onClick={() => props.onDeleteTask(task.id)}>Delete</button>
          {editState.id === task.id && (
            <button onClick={() => submitEditTaskHandler(task.id)}>Save</button>
          )}
        </li>
      ))}
    </ul>
  );
};

const ToDoApp: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);

  const addTaskHandler = (taskText: string) => {
    setTasks((prevTasks) => [
      ...prevTasks,
      { id: Math.random(), text: taskText },
    ]);
  };

  const deleteTaskHandler = (taskId: number) => {
    setTasks((prevTasks) => prevTasks.filter(task => task.id !== taskId));
  };

  const editTaskHandler = (taskId: number, newText: string) => {
    setTasks((prevTasks) => prevTasks.map(task => {
      if (task.id === taskId) {
        return {...task, text: newText};
      }
      return task;
    }));
  };

  return (
    <div>
      <TaskForm onAddTask={addTaskHandler} />
      <Task+List items={tasks} onDeleteTask={deleteTaskHandler} onEditTask={editTaskHandler} />
    </div>
  );
};

export default ToDoApp;