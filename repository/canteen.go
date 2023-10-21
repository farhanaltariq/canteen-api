package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CanteenRepository interface {
	GetAllCanteen(ctx context.Context) ([]model.Canteen, error)
	CreateCanteen(ctx context.Context, Canteen model.Canteen) error
	SeedCanteen(ctx context.Context, location string) (model.Canteen, error)
	UpdateCanteen(ctx context.Context, id int, Canteen model.Canteen) error
	DeleteCanteen(ctx context.Context, id int) error
}

type canteenRepository struct {
	db *gorm.DB
}

func NewCanteenRepository(db *gorm.DB) *canteenRepository {
	return &canteenRepository{db}
}

func (c *canteenRepository) CreateCanteen(ctx context.Context, canteen model.Canteen) error {
	err := c.db.Create(&canteen).Error
	return err // TODO: replace this
}

func (c *canteenRepository) SeedCanteen(ctx context.Context, location string) (model.Canteen, error) {
	location = strings.ReplaceAll(location, " ", "%20")
	url := "https://api.yelp.com/v3/businesses/search?location=" + location + "&sort_by=best_match&limit=1"

	req, _ := http.NewRequest("GET", url, nil)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "B_u4UEXM71BqjlNxFQdB7C17pPcV20-KVL7iGM_Qdb3mh6Qv7ED_vdR8XNb4iYug0MubTE4un0mHnAe3vrq5jIyUAx3kaihPfLHcZU7jsxqCp1pO4YNCw7psM65gZHYx"
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Canteen{}, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := model.FetchCanteenData{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
		return model.Canteen{}, err
	}

	// insert data to database

	loc := data.Businesses[0].Location.Address1 + data.Businesses[0].Location.Address2 + ", " + data.Businesses[0].Location.City
	newCanteen := model.Canteen{
		CanteenName: data.Businesses[0].Alias,
		Location:    loc,
		Contact:     data.Businesses[0].DisplayPhone,
		ImagePath:   data.Businesses[0].ImageURL,
	}
	err = c.CreateCanteen(ctx, newCanteen)

	return newCanteen, err
}

func (c *canteenRepository) GetAllCanteen(ctx context.Context) ([]model.Canteen, error) {
	canteens := []model.Canteen{}
	err := c.db.Find(&canteens).Error
	return canteens, err // TODO: replace this
}

func (c *canteenRepository) UpdateCanteen(ctx context.Context, id int, canteen model.Canteen) error {
	err := c.db.Where("id = ?", id).Updates(canteen).Error
	return err
}

func (c *canteenRepository) DeleteCanteen(ctx context.Context, id int) error {
	err := c.db.Where("id = ?", id).Delete(&model.Canteen{}).Error
	return err // TODO: replace this
}
