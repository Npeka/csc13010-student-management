# Student Management System

## Introduction

This is a **Student Management System** implemented in **Go**. The program allows users to manage a list of students with functionalities such as adding, deleting, updating, and searching students.

## Features

- **Add Student**: Input student details and save them to the list.
- **Delete Student**: Remove a student from the list using their Student ID.
- **Update Student**: Modify a student's details based on their Student ID.
- **Search Student**: Search for students by Full Name or Student ID.

## Student Information Fields

Each student record contains the following details:

- **Student ID**
- **Full Name**
- **Birth Date** (dd/mm/yyyy)
- **Gender**
- **Faculty** (e.g., Law, Business English, Japanese, French)
- **Course**
- **Program**
- **Address**
- **Email** (validated format)
- **Phone Number** (validated format)
- **Student Status** (e.g., Studying, Graduated, Withdrawn, Suspended)

---

## Folder Structure

```plaintext
student-management/
│
├── models/
│   └── student.go # Student struct definition
│
├── services/
│   └── student.service.go # Student service functions (Add, Delete, Update, Search)
│
├── tests/
│   └── student_test.go # Unit tests for student services
│
├── utils/
│   ├── to_string.go # Utility functions for string formatting
│   └── validation.go # Input validation functions (Email, Phone, etc.)
│
├── go.mod # Go module dependencies
├── main.go # Main program execution
├── Makefile # Make commands for build/run/test
├── README.md # Documentation
└── student-management.exe # Compiled executable (optional)
```

## Installation

### Prerequisites

Ensure you have **Go installed** on your system:

```sh
# Check if Go is installed
go version
```

If Go is not installed, download and install it from [Go's official website](https://go.dev/dl/).

### Install Dependencies

```sh
go mod tidy
```

---

## Usage

### Run the Program

```sh
go run main.go
```

### Build the Executable

To create an executable binary:

```sh
go build -o student-manage.exe main.go
```

### Run Tests

```sh
go test -v ./tests
```

---

## Usage Guide

When you run the program, you will see the following menu:

```sh
STUDENT MANAGEMENT SYSTEM
1. Add Student
2. Delete Student
3. Update Student
4. Search Student
5. Exit
Choose an option:
```

Follow the prompts to add, delete, update, or search for a student.

---

## Makefile (Optional)

If you have **Make** installed, you can use the provided Makefile:

```sh
# Run the program
make run

# Build the executable
make build

# Run tests
make test
```

If you don't have Make installed, you can manually run the commands as described above.
