export interface TodoStatus {
    id: string
    name: string
}

export interface Priority {
    id: string
    name: string
}

export interface User {
    id: string
    name: string
}

export interface Todo {
    id: string
    name: string
    description: string
    status: TodoStatus
    priority: Priority
    user: User
}
