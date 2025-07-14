package main

import (
	"fmt"
	"library_management/controllers"
)

func main() {
	fmt.Println("Welcome to the library management system")

	for true{
		var input int
		fmt.Println("\nMenu")
		fmt.Println("------------------")
		fmt.Println("1. List all members")
		fmt.Println("2. Add a member")
		fmt.Println("3. Remove a member")
		fmt.Println("4. List all available books")
		fmt.Println("5. List all borrowed books by member id")
		fmt.Println("6. Add a book")
		fmt.Println("7. Remove a book")
		fmt.Println("8. Borrow a book")
		fmt.Println("9. Return a book")
		fmt.Println("-1. Exit Program")
		fmt.Println("------------------")
		fmt.Println("Enter your choice: ")
		fmt.Scanf("%d\n", &input)
		switch input {
		case -1:
			fmt.Println("Exiting Program")
			return
		case 0:
			continue
		case 1:
			controllers.ListMembers()
		case 2:
			controllers.RegisterMember()
		case 3:
			controllers.DeleteMember()
		case 4:
			controllers.ListAvailableBooks()
		case 5:
			controllers.ListBorrowedBooks()
		case 6:
			controllers.RegisterBook()
		case 7:
			controllers.DeleteBook()
		case 8:
			controllers.BorrowBook()
		case 9:
			controllers.ReturnBook()
		default:
			fmt.Println("Invalid Choice")
		}
	}
}