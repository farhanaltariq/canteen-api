package service

import (
	"context"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type MenuService interface {
	CreateMenu(ctx context.Context, menu model.Menu) error
	GetAllMenu(ctx context.Context) ([]model.Menu, error)
	GetMenuByID(ctx context.Context, id int) (model.Menu, error)
	UpdateMenu(ctx context.Context, id int, menu model.Menu) error
	DeleteMenu(ctx context.Context, id int) error
}

type menuService struct {
	menuRepo repository.MenuRepository
}

func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return &menuService{menuRepo}
}

func (m *menuService) CreateMenu(ctx context.Context, menu model.Menu) error {
	err := m.menuRepo.CreateMenu(ctx, menu)
	if err != nil {
		return err
	}

	return nil
}

func (m *menuService) GetAllMenu(ctx context.Context) ([]model.Menu, error) {
	menu, err := m.menuRepo.GetAllMenu(ctx)
	return menu, err
}

func (m *menuService) GetMenuByID(ctx context.Context, id int) (model.Menu, error) {
	menu, err := m.menuRepo.GetMenuByID(ctx, id)
	return menu, err
}

func (m *menuService) UpdateMenu(ctx context.Context, id int, menu model.Menu) error {
	err := m.menuRepo.UpdateMenu(ctx, id, menu)
	return err
}

func (m *menuService) DeleteMenu(ctx context.Context, id int) error {
	err := m.menuRepo.DeleteMenu(ctx, id)
	return err
}
