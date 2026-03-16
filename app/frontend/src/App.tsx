import { useState } from 'react'
import TaskInput from "./components/TaskInput"
import TaskList from "./components/TaskList"
import './App.css'

type Task = {
  id: string
  title: string
  completed: boolean
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])

  const addTask = (title: string) => {
    const newTask: Task = {
      id: crypto.randomUUID(),
      title,
      completed: false,
    }

    setTasks([...tasks, newTask])
  }

  const toggleTask = (id: string) => {
    const updated = tasks.map((task) =>
      task.id === id
        ? { ...task, completed: !task.completed }
        : task
    )

    setTasks(updated)
  }
  return (
    <div style={{ maxWidth: "500px", margin: "40px auto" }}>
      <h1>Perfect Task</h1>

      <TaskInput onAdd={addTask} />

      <TaskList tasks={tasks} onToggle={toggleTask} />
    </div>
  )
}

export default App
