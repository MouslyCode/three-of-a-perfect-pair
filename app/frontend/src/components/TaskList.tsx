import TaskCard from "./TaskCard"

type Task = {
  id: string
  title: string
  completed: boolean
}

type Props = {
  tasks: Task[]
  onToggle: (id: string) => void
}

export default function TaskList({ tasks, onToggle }: Props) {
  return (
    <div>
      {tasks.map((task) => (
        <TaskCard key={task.id} task={task} onToggle={onToggle} />
      ))}
    </div>
  )
}