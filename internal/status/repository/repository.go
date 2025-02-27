package repository 
 
import "github.com/csc13010-student-management/internal/status" 
 
type statusRepository struct {} 
 
func NewStatusRepository() status.IStatusRepository { 
	return &statusRepository{} 
} 
