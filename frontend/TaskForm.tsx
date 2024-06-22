import React, { useState } from 'react';
import axios from 'axios';

interface TaskFormProps {
  onTaskAdded: (task: { title: string; description: string }) => void;
}

const TaskForm: React.FC<TaskFormProps> = ({ onTaskAdded }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const clearForm = () => {
    setTitle('');
    setDescription('');
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title || !description) {
      alert('Please fill in both title and description');
      return;
    }
    const newTask = {
      title,
      description,
    };
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
      </div>
      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;