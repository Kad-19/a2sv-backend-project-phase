package Domain

import "context"

const (
	CollectionUser = "users"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	FetchByEmail(c context.Context, email string) (*User, error)
}

type UserUsecase interface {
	Create(c context.Context, user *User) error
	Login(c context.Context, email string, password string) (*User, string, error)
}
