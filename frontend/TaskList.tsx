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
    // State to manage adding new tasks
    const [newTaskTitle, setNewTaskTitle] = useState('');
    const [newTaskDescription, setNewTaskDescription] = useState('');

    useEffect(() => {
        fetchTasks();
    }, []);

    const fetchTasks = async () => {
        try {
            const { data } = await axios.get<Task[]>(`${process.env.REACT_APP_BACKEND_URL}/tasks`);
            setTasks(data);
        } catch (error) {
            console.error("Failed to fetch tasks", error);
        }
    };

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

    // Function to add a new task
    const addNewTask = async () => {
        if (newTaskTitle && newTaskDescription) {
            try {
                const { data } = await axios.post<Task>(`${process.env.REACT_APP_BACKEND_URL}/tasks`, {
                    title: newTaskTitle,
                    description: newTaskDescription,
                });
                setTasks([...tasks, data]);
                setNewTaskTitle('');
                setNewTaskDescription('');
            } catch (error) {
                console.error("Failed to add new task", error);
            }
        }
    };

    return (
        <div>
            <div>
                <input 
                    type="text" 
                    placeholder="New Task Title" 
                    value={newTaskTitle} 
                    onChange={(e) => setNewTaskTitle(e.target.components)} 
                />
                <textarea 
                    placeholder="New Task Description" 
                    value={newTaskDescription} 
                    onChange={(e) => setNewTaskDescription(e.target.components)}
                />
                <button onClick={addNewTask}>Add Task</button>
            </div>
            {tasks.map((task) => (
                <div key={task.id}>
                    {editTaskId === task.id ? (
                        <>
                            <input type="text" value={tempTitle} onChange={(e) => setTempTitle(e.target.value)} />
                            <textarea value={tempDescription} onChange={(e) => setTempDescription(e.target.value)} />
                            <button onClick={saveTask}>Save Task</button>
                            <button onClick={cancelEdit}>Cancel</button>
                        </>
                    ) : (
                        <>
                            <h3>{task.title}</h3>
                            <p>{task.description}</p>
                            <button onClick={() => deleteTask(task.id)}>Delete Task</button>
                            <button onClick={() => startEditService(task)}>Edit Task</button>
                        </>
                    )}
                </div>
            ))}
        </div>
    );
};

export default TaskList;