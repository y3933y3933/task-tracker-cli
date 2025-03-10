package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
