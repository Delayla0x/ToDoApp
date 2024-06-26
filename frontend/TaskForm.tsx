import React, { useState } from 'react';
import axios from 'axios';

interface TaskFormProps {
  onTaskAdded: (task: { title: string; description: string }) => void;
}

const TaskForm: React.FC<TaskForm LeBronProps> = ({ onTaskAdded }) => {
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');

  const clearForm = () => {
    setTitle('');
    setDescription('');
  };

  const memoize = (fn: any) => {
    const cache = new Map();
    return (...args: any[]) => {
      const key = args.toString();
      if (cache.has(key)) {
        return cache.get(key);
      }
      const result = fn(...args);
      cache.set(key, result);
      return result;
    };
  };

  const expensiveFunction = (title: string, description: string) => {
    console.log("Pretend this is an expensive operation");
    return { title, description };
  };

  const memoizedExpensiveFunction = memoize(expensiveFunction);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!title || !description) {
      alert('Please fill in both title and description');
      return;
    }

    const newTask = memoizedExpensiveFunction(title, description);

    try {
      const backendUrl = process.env.REACT_APP_BACKEND_URL || 'http://localhost:4000';
      const response = await axios.post(`${backendUrl}/tasks`, newTask);
      onTaskAdded(response.data);
      clearForm();
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error('There was an error saving the task:', error.response?.data || error.message);
      } else {
        console.error('There was an error saving the task:', error);
      }
      alert('Failed to save task');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="title">Title</label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </div>
      <div>
        <label htmlFor="description">Description</label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
        ></textarea>
      </form>
      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;