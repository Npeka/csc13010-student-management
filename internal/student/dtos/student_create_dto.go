package dtos

type StudentCreateDTO struct {
	StudentID string `json:"student_id" `
	FullName  string `json:"full_name" `
	BirthDate string `json:"birth_date" `
	GenderID  int    `json:"gender_id" `
	FacultyID int    `json:"faculty_id" `
	CourseID  int    `json:"course_id" `
	ProgramID int    `json:"program_id" `
	Address   string `json:"address,omitempty" `
	Email     string `json:"email" `
	Phone     string `json:"phone" `
	StatusID  int    `json:"status_id" `
}
