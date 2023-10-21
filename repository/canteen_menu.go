package repository

import (
	"context"
	"fmt"

	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CanteenMenuRepository interface {
	GetAllCanteenMenu(ctx context.Context) ([]model.CanteenMenuDetail, error)
	GetMenuByCanteenID(ctx context.Context, id int) ([]model.Menu, error)
	CreateCanteenMenu(ctx context.Context, CanteenMenu model.CanteenMenu) error
	UpdateCanteenMenu(ctx context.Context, id int, CanteenMenu model.CanteenMenu) error
	DeleteCanteenMenu(ctx context.Context, id int) error
}

type canteenMenuRepository struct {
	db *gorm.DB
}

func NewCanteenMenuRepository(db *gorm.DB) *canteenMenuRepository {
	return &canteenMenuRepository{db}
}

func (t *canteenMenuRepository) GetAllCanteenMenu(ctx context.Context) ([]model.CanteenMenuDetail, error) {
	canteenMenu := []model.CanteenMenu{}
	res := t.db.Order("canteen_id ASC").Find(&canteenMenu)
	if res.Error != nil {
		return nil, res.Error
	}

	canteenMenuDetail := []model.CanteenMenuDetail{}

	// assign first data
	id := canteenMenu[0].CanteenID
	canteen := model.Canteen{}
	err := t.db.Where("id = ?", id).Find(&canteen).Error
	if err != nil {
		return nil, err
	}
	canteenMenuDetail = append(canteenMenuDetail, model.CanteenMenuDetail{
		CanteenName: canteen.CanteenName,
	})

	menuID := map[int]int{}
	// assign the rest
	index, count := 0, 0
	for i, v := range canteenMenu {
		if v.CanteenID == id {
			id = v.CanteenID
			count++
			menuID[index] = count
			continue
		}
		count = 1
		index++

		canteen := model.Canteen{}
		err := t.db.Where("id = ?", v.CanteenID).Find(&canteen).Error
		if err != nil {
			return nil, err
		}
		canteenMenuDetail = append(canteenMenuDetail, model.CanteenMenuDetail{
			CanteenName: canteen.CanteenName,
		})

		if i == len(canteenMenu)-1 {
			menuID[index] = count
		}
	}

	for i, len := range menuID {
		for j := 0; j < len; j++ {
			menu := model.Menu{}
			err := t.db.First(&menu, canteenMenu[j].MenuID).Error
			if err != nil {
				return nil, err
			}
			canteenMenuDetail[i].Menu = append(canteenMenuDetail[i].Menu, menu)
		}
	}

	return canteenMenuDetail, nil
}

func (t *canteenMenuRepository) CreateCanteenMenu(ctx context.Context, canteenMenu model.CanteenMenu) error {
	// check if canteen id and menu id exist
	// if not exist return error

	err := t.db.First(&model.Canteen{}, canteenMenu.CanteenID).Error
	if err != nil {
		return fmt.Errorf("canteen id %d not exist", canteenMenu.CanteenID)
	}

	err = t.db.First(&model.Menu{}, canteenMenu.MenuID).Error
	if err != nil {
		return fmt.Errorf("menu id %d not exist", canteenMenu.MenuID)
	}

	err = t.db.Create(&canteenMenu).Error
	return err // TODO: replace this
}

func (t *canteenMenuRepository) GetMenuByCanteenID(ctx context.Context, id int) ([]model.Menu, error) {
	canteenMenu := []model.CanteenMenu{}
	res := t.db.Where("canteen_id = ?", id).Find(&canteenMenu)
	if res.Error != nil {
		return nil, res.Error
	}

	menus := []model.Menu{}
	for _, v := range canteenMenu {
		menu := model.Menu{}
		res := t.db.First(&menu, v.MenuID)
		if res.Error != nil {
			return nil, res.Error
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func (t *canteenMenuRepository) UpdateCanteenMenu(ctx context.Context, id int, canteenMenu model.CanteenMenu) error {
	err := t.db.First(&model.Canteen{}, canteenMenu.CanteenID).Error
	if err != nil {
		return fmt.Errorf("canteen id %d not exist", canteenMenu.CanteenID)
	}

	err = t.db.First(&model.Menu{}, canteenMenu.MenuID).Error
	if err != nil {
		return fmt.Errorf("menu id %d not exist", canteenMenu.MenuID)
	}

	res := t.db.First(&model.CanteenMenu{}, id).Updates(&canteenMenu)
	if res.RowsAffected == 0 {
		return res.Error
	}

	return nil
}

func (t *canteenMenuRepository) DeleteCanteenMenu(ctx context.Context, id int) error {
	res := t.db.Delete(&model.CanteenMenu{}, id)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
