package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Subject struct {
	name  string
	grade float64
}

func main() {
	var studentName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName)
	if studentName == "" {
		fmt.Println("Error: Name cannot be empty")
		return
	}

	var numSubjects int
	fmt.Print("Enter the number of subjects: ")
	_, err := fmt.Scanf("%d\n", &numSubjects)
	if err != nil || numSubjects <= 0 {
		fmt.Println("Error: Please enter a valid positive number of subjects")
		return
	}

	subjects := make([]Subject, 0, numSubjects)

	for i := 0; i < numSubjects; i++ {
		var subjectName string
		fmt.Printf("Enter subject %d name: ", i+1)
		fmt.Scanln(&subjectName)
		subjectName = strings.TrimSpace(subjectName)
		if subjectName == "" {
			fmt.Println("Error: Subject name cannot be empty")
			return
		}

		var gradeStr string
		fmt.Printf("Enter grade for %s (0-100): ", subjectName)
		fmt.Scanln(&gradeStr)
		grade, err := strconv.ParseFloat(strings.TrimSpace(gradeStr), 64)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("Error: Please enter a valid grade between 0 and 100")
			return
		}

		subjects = append(subjects, Subject{name: subjectName, grade: grade})
	}

	average := calculateAverage(subjects)
	displayResults(studentName, subjects, average)
}

func calculateAverage(subjects []Subject) float64 {
	if len(subjects) == 0 {
		return 0
	}
	var sum float64
	for _, subject := range subjects {
		sum += subject.grade
	}
	return sum / float64(len(subjects))
}

func displayResults(studentName string, subjects []Subject, average float64) {
	fmt.Printf("\nGrade Report for %s\n", studentName)
	fmt.Println("-----------------------------")
	for _, subject := range subjects {
		fmt.Printf("%s: %.2f\n", subject.name, subject.grade)
	}
	fmt.Printf("-----------------------------")
	fmt.Printf("\nAverage Grade: %.2f\n", average)
}
