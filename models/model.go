package models



import (
    "github.com/google/uuid"
    "time"
)

type User struct {
    UserID   uuid.UUID
    FullName string
    Username string
    Password string
    Gmail    string
}

type Todo struct {
    TodoID      uuid.UUID
    Task        string
    CreatedAt   time.Time
    IsCompleted bool
}


type GetTodosResponse struct {
	Todos []Todo
	Count int
}