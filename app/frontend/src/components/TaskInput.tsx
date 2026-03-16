import { useState } from "react";

type Props = {
    onAdd: (title: string) => void
}

export default function TaskInput({ onAdd }: Props) {
    const [title, setTitle] = useState("")

    const handleSubmit = () => {
        if (!title.trim()) return
        onAdd(title)
        setTitle("")
    }

    return (
        <div style={{ display: "flex", gap: "8px", marginBottom: "20px" }}>
            <input
                type="text"
                placeholder="Add task..."
                value={title}
                onChange={(e) => setTitle(e.target.value)}
            />
            <button onClick={handleSubmit}>Add</button>
        </div>
    )
}

