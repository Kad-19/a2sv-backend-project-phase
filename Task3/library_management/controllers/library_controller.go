package controllers

import (
	"errors"
	"os"
	"fmt"
	"library_management/services"
	"library_management/models"
	"github.com/olekukonko/tablewriter"
)
var lms services.LibraryManager
var library services.Library = &lms
var booksCount, memberCount int

func TakeStrInput(prompt string) (string, error) {
	var input string;

	fmt.Println(prompt)
	fmt.Scanln(&input)
	for input == "" {
		fmt.Println("\nInput is required, please re enter (enter << to cancel request): ")
		fmt.Scanln(&input)
		if input == "<<" {
			return input, errors.New("\nrequest canceled")
		}
	}
	return input, nil
}

func TakeIntInput(prompt string) (int, error) {
	var input int;

	fmt.Println(prompt)
	fmt.Scanf("%d\n", &input)
	for input == 0 {
		fmt.Println("\nInput is required, please re enter (enter 0 to cancel request): ")
		fmt.Scanln(&input)
		if input == 0 {
			return input, errors.New("\nrequest canceled")
		}
	}
	return input, nil

}

func PrintBooks(books []models.Book) {
	table := tablewriter.NewWriter(os.Stdout)

    table.SetHeader([]string{"ID", "Title", "Author", "Status"})

    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    table.SetCenterSeparator("|")
    table.SetColumnSeparator("|")
    table.SetRowSeparator("-")

    for _, book := range books {
        table.Append([]string{
            fmt.Sprintf("%d", book.Id),
            book.Title,
            book.Author,
			string(book.Status),
        })
    }

    // Render the table
    table.Render()
}
func PrintMembers(members []models.Member) {
	table := tablewriter.NewWriter(os.Stdout)

    table.SetHeader([]string{"ID", "Name"})

    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    table.SetCenterSeparator("|")
    table.SetColumnSeparator("|")
    table.SetRowSeparator("-")

    for _, member := range members {
        table.Append([]string{
            fmt.Sprintf("%d", member.Id),
            member.Name,
        })
    }

    // Render the table
    table.Render()
}
func RegisterBook() {
	title, err := TakeStrInput("\nEnter the title of the book: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	author, err := TakeStrInput("\nEnter the author of the book: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	booksCount += 1
	book := models.Book{Id: booksCount ,Author: author, Title: title, Status: "Available"}

	library.AddBook(book)
	fmt.Printf("\n Book successfuly added")

}
func RegisterMember() {
	name, err := TakeStrInput("\nEnter Name: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	memberCount += 1
	member := models.Member{Id: memberCount, Name: name}

	library.AddMember(member)
	fmt.Printf("\n Member successfuly add")

}

func DeleteBook() {
	bookId, err := TakeIntInput("\nEnter id of the book: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	e := library.RemoveBook(bookId)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("\n Book %d successfuly deleted", bookId)

}
func DeleteMember() {
	memberId, err := TakeIntInput("\nEnter id of the member: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	e := library.RemoveMember(memberId)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("\n Member %d successfuly deleted", memberId)

}

func BorrowBook() {
	bookId, err := TakeIntInput("\nEnter id of the book: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	memberId, err := TakeIntInput("\nEnter id of the member: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	e := library.BorrowBook(bookId, memberId)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("\nSuccesfully borrowed book %d for %d", bookId, memberId)
}
func ReturnBook() {
	bookId, err := TakeIntInput("\nEnter id of the book: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	memberId, err := TakeIntInput("\nEnter id of the member: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	e := library.ReturnBook(bookId, memberId)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("\nSuccesfully returned book %d for %d", bookId, memberId)

}

func ListAvailableBooks() {
	books := library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No Available books found")
		return
	}
	PrintBooks(books)
}

func ListMembers() {
	members := library.ListMembers()
	if len(members) == 0 {
		fmt.Println("No members found")
		return
	}
	PrintMembers(members)
}
func ListBorrowedBooks() {
	memberId, err := TakeIntInput("\nEnter id of the member: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	books, e := library.ListBorrowedBooks(memberId)
	if e != nil {
		fmt.Println(e)
		return
	}

	if len(books) == 0 {
		fmt.Println("No Borrows books found")
		return
	}
	PrintBooks(books)
}