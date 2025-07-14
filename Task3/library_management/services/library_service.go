package services

import (
	"errors"
	"library_management/models"
)
type Library interface {
	AddBook(book models.Book)
	RemoveBook(bookId int) error
	AddMember(member models.Member)
	RemoveMember(memberId int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
	ListMembers() []models.Member
}

type LibraryManager struct {
	books map[int]models.Book;
	members map[int]models.Member
}


func (lm *LibraryManager) AddBook(book models.Book) {
	if lm.books == nil {
		lm.books = make(map[int]models.Book)
	}
	lm.books[book.Id] = book
}

func (lm *LibraryManager) RemoveBook(bookId int) error{
	b, ok := lm.books[bookId]
	if ok {
		if b.Status == "Available" {
			delete(lm.books, bookId)
			return nil
		}
		return errors.New("book is on loan. please return book before attempting to delete")
	}
	return errors.New("book not found")
}
func (lm *LibraryManager) AddMember(member models.Member) {
	if lm.members == nil {
		lm.members = make(map[int]models.Member)
	}
	lm.members[member.Id] = member
}

func (lm *LibraryManager) RemoveMember(memberId int) error{
	m, ok := lm.members[memberId]
	if ok {
		if len(m.Borrowedbooks) == 0 {
			delete(lm.members, memberId)
			return nil
		}
		return errors.New("member had borrowed books. please return book before attempting to delete")
	}
	return errors.New("member not found")
}
func (lm *LibraryManager) BorrowBook(bookID int, memberID int) error{

	book, ok := lm.books[bookID]
	if ok {
		if book.Status == "Borrowed" {
			return errors.New("this book is not available")
		}
		member, mok := lm.members[memberID]

		if mok {
			book.Status = "Borrowed"
			member.Borrowedbooks = append(member.Borrowedbooks, book)
			lm.members[memberID] = member
			lm.books[bookID] = book
		} else {
			return errors.New("this member does not exist")
		}

		
	} else {
		return errors.New("this book does not exist")
	}

	return nil

}
func (lm *LibraryManager) ReturnBook(bookID int, memberID int) error{
	book, ok := lm.books[bookID]
	if ok {
		if book.Status == "Available" {
			return errors.New("this book is not borrowed")
		}
		member, mok := lm.members[memberID]

		if mok {
			borrowedbooks, err := RemoveBookByID(member.Borrowedbooks, bookID)
			if err != nil {
				return err
			}
			member.Borrowedbooks = borrowedbooks
			lm.members[memberID] = member
			book.Status = "Available"
			lm.books[bookID] = book
		} else {
			return errors.New("this member does not exist")
		}

		
	} else {
		return errors.New("this book does not exist")
	}

	return nil

}
func (lm *LibraryManager) ListAvailableBooks() []models.Book{
	var	availableBooks []models.Book
	for _, b := range lm.books {
		if b.Status == "Available" {
			availableBooks = append(availableBooks, b)
		}
	}
	return availableBooks
}
func (lm *LibraryManager) ListBorrowedBooks(memberID int) ([]models.Book, error){
	var	borrowedBooks []models.Book
	member, ok := lm.members[memberID]
	if ok {
		borrowedBooks = append(borrowedBooks, member.Borrowedbooks...)

	} else {
		return borrowedBooks, errors.New("member not found")
	}
	return borrowedBooks, nil
}

func (lm *LibraryManager) ListMembers() []models.Member{
	var	members []models.Member
	for _, m := range lm.members {
		members = append(members, m)
	}
	return members
}

func RemoveBookByID(books []models.Book, id int) ([]models.Book, error) {
    for i, book := range books {
        if book.Id == id {
            
            return append(books[:i], books[i+1:]...), nil
        }
    }
    return books, errors.New("book not found")
}
