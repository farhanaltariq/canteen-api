package service

import (
	"context"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type CanteenMenuService interface {
	CreateCanteenMenu(ctx context.Context, menu model.CanteenMenu) error
	GetAllCanteenMenu(ctx context.Context) ([]model.CanteenMenuDetail, error)
	GetMenuByCanteenID(ctx context.Context, id int) ([]model.Menu, error)
	UpdateCanteenMenu(ctx context.Context, id int, menu model.CanteenMenu) error
	DeleteCanteenMenu(ctx context.Context, id int) error
}

type canteenMenuService struct {
	canteenMenuRepo repository.CanteenMenuRepository
}

func NewCanteenMenuService(canteenMenuRepo repository.CanteenMenuRepository) CanteenMenuService {
	return &canteenMenuService{canteenMenuRepo}
}

func (cm *canteenMenuService) CreateCanteenMenu(ctx context.Context, canteenMenu model.CanteenMenu) error {
	err := cm.canteenMenuRepo.CreateCanteenMenu(ctx, canteenMenu)
	if err != nil {
		return err
	}

	return nil
}

func (cm *canteenMenuService) GetAllCanteenMenu(ctx context.Context) ([]model.CanteenMenuDetail, error) {
	canteenMenu, err := cm.canteenMenuRepo.GetAllCanteenMenu(ctx)
	return canteenMenu, err
}
func (cm *canteenMenuService) GetMenuByCanteenID(ctx context.Context, id int) ([]model.Menu, error) {
	canteenMenu, err := cm.canteenMenuRepo.GetMenuByCanteenID(ctx, id)
	return canteenMenu, err
}
func (cm *canteenMenuService) UpdateCanteenMenu(ctx context.Context, id int, canteenMenu model.CanteenMenu) error {
	err := cm.canteenMenuRepo.UpdateCanteenMenu(ctx, id, canteenMenu)
	if err != nil {
		return err
	}

	return nil
}

func (cm *canteenMenuService) DeleteCanteenMenu(ctx context.Context, id int) error {
	err := cm.canteenMenuRepo.DeleteCanteenMenu(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
