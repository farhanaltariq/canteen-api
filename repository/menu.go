package repository

import (
	"context"

	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAllMenu(ctx context.Context) ([]model.Menu, error)
	GetMenuByID(ctx context.Context, id int) (model.Menu, error)
	CreateMenu(ctx context.Context, Menu model.Menu) error
	UpdateMenu(ctx context.Context, id int, Menu model.Menu) error
	DeleteMenu(ctx context.Context, id int) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db}
}

func (m *menuRepository) CreateMenu(ctx context.Context, menu model.Menu) error {
	err := m.db.Create(&menu).Error
	return err
}

func (m *menuRepository) GetAllMenu(ctx context.Context) ([]model.Menu, error) {
	menu := []model.Menu{}
	err := m.db.Find(&menu).Error
	return menu, err
}

func (m *menuRepository) GetMenuByID(ctx context.Context, id int) (model.Menu, error) {
	menu := model.Menu{}
	err := m.db.Find(&menu, id).Error
	return menu, err
}

func (m *menuRepository) UpdateMenu(ctx context.Context, id int, menu model.Menu) error {
	err := m.db.First(&menu, id).Error
	if err != nil {
		return err
	}
	err = m.db.Model(&model.Menu{}).Where("id = ?", id).Updates(&menu).Error
	return err
}

func (m *menuRepository) DeleteMenu(ctx context.Context, id int) error {
	err := m.db.First(&model.Menu{}, id).Delete(&model.Menu{}).Error
	return err
}
