package service

import (
	"context"
	"errors"
	"time"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type UserService interface {
	Login(ctx context.Context, user *model.User) (id int, err error)
	Register(ctx context.Context, user *model.User) (model.User, error)

	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(ctx context.Context, user *model.User) (id int, err error) {
	dbUser, err := s.userRepository.GetUserByusername(ctx, user.Username)
	if err != nil {
		return 0, err
	}

	if dbUser.Username == "" || dbUser.ID == 0 {
		return 0, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return 0, errors.New("wrong username or password")
	}

	return dbUser.ID, nil
}

func (s *userService) Register(ctx context.Context, user *model.User) (model.User, error) {
	dbUser, err := s.userRepository.GetUserByusername(ctx, user.Username)
	if err != nil {
		return *user, err
	}

	if dbUser.Username != "" || dbUser.ID != 0 {
		return *user, errors.New("username already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepository.CreateUser(ctx, *user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := s.userRepository.GetUserByusername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
