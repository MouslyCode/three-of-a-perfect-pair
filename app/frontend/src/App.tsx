import { useState, useEffect } from 'react'
import TaskInput from "./components/TaskInput"
import TaskList from "./components/TaskList"
import { getTasks, createTask, toggleTask } from './services/api'
import './App.css'

type Task = {
  id: string
  title: string
  completed: boolean
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])

  useEffect(() => {
    fetchTasks()
  }, [])

  const fetchTasks = async () => {
    const data = await getTasks()
    setTasks(data)
  }

  const addTask = async (title: string) => {
    await createTask(title)
    fetchTasks()
  }

  const handleToggle = async (id: string) => {
    await toggleTask(id)
    fetchTasks()
  }
  return (
    <div style={{ maxWidth: "500px", margin: "40px auto" }}>
      <h1>Task List</h1>

      <TaskInput onAdd={addTask} />

      <TaskList tasks={tasks} onToggle={handleToggle} />
    </div>
  )
}

export default App
