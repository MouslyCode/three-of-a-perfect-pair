const BASE_URL = "http://localhost:8080"

export async function getTasks() {
    const res = await fetch(`${BASE_URL}/tasks`)
    return res.json()
}

export async function createTask(title: string) {
    const res = await fetch(`${BASE_URL}/tasks`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ title }),
    })
    return res.json()
}

export async function toggleTask(id: string) {
    const res = await fetch(`${BASE_URL}/tasks/${id}`, {
        method: "PATCH",
    })
    return res.json()
}

export async function deleteTask(id: string) {
    const res = await fetch(`${BASE_URL}/tasks/${id}`, {
        method: "DELETE",
    })
    return res.json()
}