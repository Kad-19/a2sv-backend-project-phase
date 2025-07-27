package usecase

import (
	"context"
	"errors"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	"time"
)

type userUsecase struct {
	userRepo      Domain.UserRepository
	contextTimeout time.Duration
	env           *Infrastructure.Env
}

func NewUserUsecase(userRepo Domain.UserRepository, timeout time.Duration, env *Infrastructure.Env) *userUsecase {
	return &userUsecase{
		userRepo:      userRepo,
		contextTimeout: timeout,
		env:           env,
	}
}

func (uu *userUsecase) Create(c context.Context, user *Domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	// Check if user already exists
	existingUser, err := uu.userRepo.FetchByEmail(ctx, user.Email)
	if err != nil && existingUser != nil {
		return errors.New("user already exists")
	}
	// Hash the password before saving
	hashedPassword, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return uu.userRepo.Create(ctx, user)
}

func (uu *userUsecase) Login(c context.Context, email string, password string) (*Domain.User, string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepo.FetchByEmail(ctx, email)
	if err != nil || Infrastructure.CheckPasswordHash(user.Password, password) != nil {
		return nil, "", errors.New("invalid email or password")
	}
	// Generate JWT token
	token, err := Infrastructure.GenerateJWTToken(user, uu.env)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
