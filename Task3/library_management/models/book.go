package models

type Status string

const (
	Borrowed  Status = "borrowed"
	Available Status = "available"
)

type Book struct {
	Id int;
	Title string;
	Author string;
	Status Status;
}