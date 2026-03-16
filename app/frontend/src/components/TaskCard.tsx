type Task = {
    id: string
    title: string
    completed: boolean
}

type Props = {
    task: Task
    onToggle: (id: string) => void
}

export default function TaskCard({ task, onToggle }: Props) {
    return (
        <div
            style={{
                display: "flex",
                alignItems: "center",
                gap: "10px",
                padding: "10px",
                border: "1px solid #ddd",
                borderRadius: "8px",
                marginBottom: "10px",
            }}
        >
            <input
                type="checkbox"
                checked={task.completed}
                onChange={() => onToggle(task.id)}
            />

            <span
                style={{
                    textDecoration: task.completed ? "line-through" : "none",
                }}
            >
                {task.title}
            </span>
        </div>
    )
}