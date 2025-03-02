package events

type UserCreatedEvent struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}
