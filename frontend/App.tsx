import React, { useState, useEffect } from 'react';

interface Task {
  id: number;
  text: string;
}

const TaskForm: React.FC<{ onAddTask: (taskText: string) => void }> = (props) => {
  const [enteredText, setEnteredText] = useState("");

  const addTaskHandler = (event: React.FormEvent) => {
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

const TaskList: React.FC<{ items: Task[]; onDeleteTask: (taskId: number) => void }> = (props) => {
  return (
    <ul>
      {props.items.map(task => (
        <li key={task.id}>
          {task.text}
          <button onClick={() => props.onDeleteTask(task.id)}>Delete</button>
        </li>
      ))}
    </ul>
  );
};

const ToDoApp: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);

  const addTaskHandler = (taskText: string) => {
    setTasks((prevTasks) => [
      ...prevGroups,
      { id: Math.random(), text: taskText },
    ]);
  };

  const deleteTaskHandler = (taskId: number) => {
    setTasks((prevTasks) => prevTasks.filter(task => task.id !== taskId));
  };

  return (
    <div>
      <TaskForm onAddTask={addTaskHandler} />
      <TaskList items={tasks} onDeleteTask={deleteTaskHandler} />
    </div>
  );
};

export default ToDoApp;