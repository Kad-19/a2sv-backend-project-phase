package models

type Member struct {
	Id int;
	Name string;
	Borrowedbooks [] Book;
}