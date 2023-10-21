package service

import (
	"context"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type CanteenService interface {
	CreateCanteen(ctx context.Context, canteen model.Canteen) error
	SeedCanteen(ctx context.Context, location string) (model.Canteen, error)
	GetAllCanteen(ctx context.Context) ([]model.Canteen, error)
	UpdateCanteen(ctx context.Context, id int, canteen model.Canteen) error
	DeleteCanteen(ctx context.Context, id int) error
}

type canteenService struct {
	canteenRepo repository.CanteenRepository
	userRepo    repository.UserRepository
}

func NewCanteenService(canteenRepo repository.CanteenRepository, userRepo repository.UserRepository) CanteenService {
	return &canteenService{canteenRepo, userRepo}
}

func (t *canteenService) CreateCanteen(ctx context.Context, canteen model.Canteen) error {
	err := t.canteenRepo.CreateCanteen(ctx, canteen)
	return err
}

func (t *canteenService) SeedCanteen(ctx context.Context, location string) (model.Canteen, error) {
	res, err := t.canteenRepo.SeedCanteen(ctx, location)
	return res, err
}

func (t *canteenService) GetAllCanteen(ctx context.Context) ([]model.Canteen, error) {
	canteen, err := t.canteenRepo.GetAllCanteen(ctx)
	return canteen, err
}
func (t *canteenService) UpdateCanteen(ctx context.Context, id int, canteen model.Canteen) error {
	err := t.canteenRepo.UpdateCanteen(ctx, id, canteen)
	return err
}

func (t *canteenService) DeleteCanteen(ctx context.Context, id int) error {
	err := t.canteenRepo.DeleteCanteen(ctx, id)
	return err
}
