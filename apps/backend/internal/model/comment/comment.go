package comment

import (
	"github.com/google/uuid"
	"github.com/sriniously/go-tasker/internal/model"
)

type Comment struct {
	model.Base
	UserID  uuid.UUID `json:"userId" db:"user_id"`
	Content string    `json:"content" db:"content"`
	TodoID uuid.UUID 	`json:"todoId" db:"todo_id"`
}
