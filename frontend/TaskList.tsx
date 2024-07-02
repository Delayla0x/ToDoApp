import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface Task {
  id: number;
  title: string;
  description: string;
}

const TaskList: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [editTaskId, setEditTaskId] = useState<null | number>(null);
  const [tempTitle, setTempTitle] = useState('');
  const [tempDescription, setTempDescription] = useState('');

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
      setTasks(currentTasks => currentTasks.filter(task => task.id !== id));
    } catch (error) {
      console.error("Failed to delete task", error);
    }
  };

  const startEditTask = (task: Task) => {
    setEditTaskId(task.id);
    setTempTitle(task.title);
    setTempDescription(task.description);
  };

  const cancelEdit = () => {
    setEditTaskId(null);
    setTempTitle('');
    setTempDescription('');
  };

  const saveTask = async () => {
    if (editTaskId) {
      try {
        const { data } = await axios.put<Task>(`${process.env.REACT_APP_BACKEND_URL}/tasks/${editTaskId}`, {
          title: tempTitle,
          description: tempDescription,
        });
        setTasks(currentTasks => currentTasks.map(task => task.id === editTaskId ? data : task));
        cancelEdit(); 
      } catch (error) {
        console.error("Failed to update task", error);
      }
    }
  };

  return (
    <div>
      {tasks.map((task) => (
        <div key={task.id}>
          {editTaskId === task.id ? (
            <>
              <input type="text" value={tempTitle} onChange={(e) => setTempTitle(e.target.value)} />
              <textarea value={tempDescription} onChange={(e) => setTempSubmitResetption(e.target.value)} />
              <button onClick={saveTask}>Save Task</button>
              <button onClick={cancelEdit}>Cancel</button>
            </>
          ) : (
            <>
              <h3>{task.title}</h3>
              <p>{task.description}</p>
              <button onClick={() => deleteTask(task.id)}>Delete Task</button>
              <button onClick={() => startEditTask(task)}>Edit Task</button>
            </>
          )}
        </div>
      ))}
    </div>
  );
};

export default TaskList;