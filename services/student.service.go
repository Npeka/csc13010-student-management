package services

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"student-management/models"
	"student-management/utils"
)

var Students []models.Student

func IsStudentIDExists(id string) bool {
	for _, s := range Students {
		if s.ID == id {
			return true
		}
	}
	return false
}

func AddStudent(in io.Reader, out io.Writer) {
	var s models.Student
	reader := bufio.NewReader(in)

	fmt.Fprint(out, "Enter Student ID: ")
	s.ID, _ = reader.ReadString('\n')
	s.ID = strings.TrimSpace(s.ID)

	if IsStudentIDExists(s.ID) {
		fmt.Fprintln(out, "Student ID already exists! Please enter a unique ID.")
		return
	}

	fmt.Fprint(out, "Enter Full Name: ")
	s.FullName, _ = reader.ReadString('\n')
	s.FullName = strings.TrimSpace(s.FullName)

	fmt.Fprint(out, "Enter Birth Date (dd/mm/yyyy): ")
	s.BirthDate, _ = reader.ReadString('\n')
	s.BirthDate = strings.TrimSpace(s.BirthDate)

	fmt.Fprint(out, "Enter Gender: ")
	s.Gender, _ = reader.ReadString('\n')
	s.Gender = strings.TrimSpace(s.Gender)

	fmt.Fprintf(out, "Enter Faculty (%v): ", utils.ArrayToString(utils.ValidFaculties))
	s.Faculty, _ = reader.ReadString('\n')
	s.Faculty = strings.TrimSpace(s.Faculty)
	if !utils.IsValidFaculty(s.Faculty) {
		fmt.Fprintln(out, "Invalid Faculty!", s.Faculty)
		return
	}

	fmt.Fprint(out, "Enter Course: ")
	s.Course, _ = reader.ReadString('\n')
	s.Course = strings.TrimSpace(s.Course)

	fmt.Fprint(out, "Enter Program: ")
	s.Program, _ = reader.ReadString('\n')
	s.Program = strings.TrimSpace(s.Program)

	fmt.Fprint(out, "Enter Address: ")
	s.Address, _ = reader.ReadString('\n')
	s.Address = strings.TrimSpace(s.Address)

	fmt.Fprint(out, "Enter Email: ")
	s.Email, _ = reader.ReadString('\n')
	s.Email = strings.TrimSpace(s.Email)
	if !utils.IsValidEmail(s.Email) {
		fmt.Fprintln(out, "Invalid Email!")
		return
	}

	fmt.Fprint(out, "Enter Phone Number: ")
	s.Phone, _ = reader.ReadString('\n')
	s.Phone = strings.TrimSpace(s.Phone)
	if !utils.IsValidPhone(s.Phone) {
		fmt.Fprintln(out, "Invalid Phone Number!")
		return
	}

	fmt.Fprintf(out, "Enter Student Status (%v): ", utils.ArrayToString(utils.ValidStatuses))
	s.Status, _ = reader.ReadString('\n')
	s.Status = strings.TrimSpace(s.Status)
	if !utils.IsValidStatus(s.Status) {
		fmt.Fprintln(out, "Invalid Status!")
		return
	}

	Students = append(Students, s)
	fmt.Fprintln(out, "Student added successfully!")
}

func DeleteStudent(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	fmt.Fprint(out, "Enter Student ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for i, s := range Students {
		if s.ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			fmt.Fprintln(out, "Student deleted successfully!")
			return
		}
	}
	fmt.Fprintln(out, "Student ID not found.")
}

func UpdateStudent(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	fmt.Fprint(out, "Enter Student ID to update: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for i, s := range Students {
		if s.ID == id {
			fmt.Fprint(out, "Enter New Email: ")
			s.Email, _ = reader.ReadString('\n')
			s.Email = strings.TrimSpace(s.Email)
			if !utils.IsValidEmail(s.Email) {
				fmt.Fprintln(out, "Invalid Email!")
				return
			}

			fmt.Fprint(out, "Enter New Phone Number: ")
			s.Phone, _ = reader.ReadString('\n')
			s.Phone = strings.TrimSpace(s.Phone)
			if !utils.IsValidPhone(s.Phone) {
				fmt.Fprintln(out, "Invalid Phone Number!")
				return
			}

			fmt.Fprintf(out, "Enter New Status (%v): ", utils.ArrayToString(utils.ValidStatuses))
			s.Status, _ = reader.ReadString('\n')
			s.Status = strings.TrimSpace(s.Status)
			if !utils.IsValidStatus(s.Status) {
				fmt.Fprintln(out, "Invalid Status!")
				return
			}

			Students[i] = s
			fmt.Fprintln(out, "Student updated successfully!")
			return
		}
	}
	fmt.Fprintln(out, "Student ID not found.")
}

func SearchStudent(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	fmt.Fprint(out, "Enter Student ID or Full Name to search: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)

	found := false
	for _, s := range Students {
		if strings.Contains(s.ID, keyword) || strings.Contains(s.FullName, keyword) {
			fmt.Fprintf(out, "ID: %s, Name: %s, Email: %s, Phone: %s, Status: %s\n",
				s.ID, s.FullName, s.Email, s.Phone, s.Status)
			found = true
		}
	}
	if !found {
		fmt.Fprintln(out, "No student found.")
	}
}
