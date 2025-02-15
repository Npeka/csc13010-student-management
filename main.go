package main

import (
	"bufio"
	"fmt"
	"os"
	"student-management/services"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSTUDENT MANAGEMENT SYSTEM")
		fmt.Println("1. Add Student")
		fmt.Println("2. Delete Student")
		fmt.Println("3. Update Student")
		fmt.Println("4. Search Student")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {
		case 1:
			services.AddStudent(os.Stdin, os.Stdout)
		case 2:
			services.DeleteStudent(os.Stdin, os.Stdout)
		case 3:
			services.UpdateStudent(os.Stdin, os.Stdout)
		case 4:
			services.SearchStudent(os.Stdin, os.Stdout)
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
