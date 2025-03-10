package dtos

type StudentResponseDTO struct {
	ID        uint   `json:"id"`
	StudentID string `json:"student_id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Faculty   string `json:"faculty"`
	Course    string `json:"course"`
	Program   string `json:"program"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
}
