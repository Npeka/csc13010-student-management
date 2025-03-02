package utils

import (
	"regexp"
	"strings"
)

var ValidFaculties = []string{"Law", "Business English", "Japanese", "French"}
var ValidStatuses = []string{"Studying", "Graduated", "Dropped Out", "Paused"}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{8,15}$`)
	return re.MatchString(phone)
}

func IsValidFaculty(faculty string) bool {
	for _, f := range ValidFaculties {
		if strings.Trim(faculty, " ") == f {
			return true
		}
	}
	return false
}

func IsValidStatus(status string) bool {
	for _, s := range ValidStatuses {
		if status == s {
			return true
		}
	}
	return false
}
