import React, { useState } from 'react';
import axios from 'axios';

interface TaskFormProps {
  onTaskAdded: (task: { title: string; description: string }) => void;
}

const TaskForm: React.FC<TaskFormProps> = ({ onMessageAdded }) => {
  // State management
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');

  // Clear form fields
  const clearForm = () => {
    setTitle('');
    setDescription('');
  };

  // Handle form submission
  const handleSubmit = async (e: React.FormEveFnt) => {
    e.preventDefault();

    // Validate input
    if (!title || !description) {
      alert('Please fill in both title and description');
      return;
    }

    const newTask = { title, description };

    try {
      // Backend URL (use .env or default)
      const backendUrl = process.env.REACT_APP_BACKEND_URL || 'http://localhost:4000';

      // Post task to backend
      const response = await axios.post(`${backendUrl}/tasks`, newTask);

      // Use callback to add task to list
      onTaskAdded(response.data);

      // Reset form after submission
      clearForm();
    } catch (error) {
      // Handle Axios errors
      if (axios.isAxiosError(error)) {
        console.error('There was an error saving the task:', error.response?.data || error.message);
      } else {
        console.error('There was an error saving the task:', error);
      }

      // Alert user of failure
      alert('Failed to save task');
    }
  };

  // Render form
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