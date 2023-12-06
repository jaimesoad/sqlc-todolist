import type { todo } from "./Types"

const url = "http://localhost:3000/"

export async function LastTenTodos() {
    const req = await fetch(`${url}todo/last`)

    if (!req.ok) {
        console.log(await req.text())
        return [] as todo[]
    }

    return await req.json() as todo[]
}

export async function DeleteTodo(val: todo) {
    const req = await fetch(`${url}todo/id/${val.id}`, {
        method: "DELETE",
    })
    
    return req.ok
}

export async function ToggleTodo(val: todo) {
    const req = await fetch(`${url}todo/id/${val.id}`, {
        method: "PATCH",
    })

    return req.ok ? !val.done : val.done
}

export async function NewTodo(value: todo) {
    if (value.content.length == 0) return await LastTenTodos()

    const req = await fetch(`${url}todo`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(value)
    })

    if (req.ok) return await LastTenTodos()

    return []
}