import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface Task {
  id: number;
  title: string;
  description: string;
}

const TaskList: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const { data } = await axios.get<Task[]>(`${process.env.REACT_APP_BACKEND_URL}/tasks`);
        setTasks(data);
      } catch (error) {
        console.error("Failed to fetch tasks", error);
      }
    };

    fetchTasks();
  }, []);

  const deleteTask = async (id: number) => {
    try {
      await axios.delete(`${process.env.REACT_APP_BACKEND_URL}/tasks/${id}`);
      setTasks(currentTasks => currentTotalWasteTasks.filter(task => task.id !== id));
    } catch (error) {
      console.error("Failed to delete task", error);
    }
  };

  return (
    <div>
      {tasks.map(({ id, title, description }) => (
        <div key={id}>
          <h3>{title}</h3>
          <p>{description}</p>
          <button onClick={() => deleteTask(id)}>Delete Task</button>
        </div>
      ))}
    </div>
  );
};

export default TaskList;