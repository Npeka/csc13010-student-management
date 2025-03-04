package events

import "github.com/google/uuid"

type StudentCreatedEvent struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}
