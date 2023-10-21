package repository

import (
	"context"

	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetUserByusername(ctx context.Context, username string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil
}

func (r *userRepository) GetUserByusername(ctx context.Context, username string) (model.User, error) {
	user := model.User{}
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	// get row length
	var count int64
	r.db.WithContext(ctx).Table("users").Count(&count)
	user.ID = int(count) + 1

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
