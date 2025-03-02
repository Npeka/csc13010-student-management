package events

type StudentCreatedEvent struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}
