package events

import "github.com/google/uuid"

type UserCreatedEvent struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}
