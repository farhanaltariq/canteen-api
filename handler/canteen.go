package api

import (
	"fmt"
	"net/http"
	"strconv"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"

	"github.com/gin-gonic/gin"
)

type CanteenAPI interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Seed(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type canteenAPI struct {
	canteenService service.CanteenService
}

func NewCanteenAPI(canteenService service.CanteenService) *canteenAPI {
	return &canteenAPI{canteenService}
}

//	Create Canteen
//
// @Summary Create Canteen
// @Description Create Canteen
// @Tags Canteen
// @Security Authorization
// @Accept json
// @Produce json
// @Param Canteen body model.CanteenData true "Canteen"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /canteens [post]
func (t *canteenAPI) Create(c *gin.Context) {
	canteen := model.Canteen{}
	if err := c.BindJSON(&canteen); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	// cookies, _ := c.Request.Cookie("session_token")
	userID, _ := c.Get("id")
	intUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	err := t.canteenService.CreateCanteen(c.Request.Context(), canteen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	c.JSON(http.StatusOK, model.NewSuccessResponse(intUserID, "canteen created"))
	// TODO: answer here
}

//	Seed Canteen
//
// @Summary Seed Canteen Database
// @Description Adding 1 Canteen to Database from Third Prty API (YELP API)
// @Description
// @Tags Canteen
// @Security Authorization
// @Produce json
// @Param location path string true "Canteen location (ex. : New York, Singapore)"
// @Success 201 {object} model.Canteen
// @Failure 400 {object} model.ErrorResponse
// @Router /canteens/seed/{location} [get]
func (t *canteenAPI) Seed(c *gin.Context) {
	location := c.Params.ByName("location")
	canteen, err := t.canteenService.SeedCanteen(c.Request.Context(), location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	c.JSON(http.StatusCreated, canteen)
}

//	Get all canteen
//
// @Summary Get all canteen
// @Description Get all canteen
// @Tags Canteen
// @Produce json
// @Success 200 {object} []model.Canteen
// @Failure 400 {object} model.ErrorResponse
// @Router /canteens [get]
func (t *canteenAPI) Get(c *gin.Context) {
	canteens, err := t.canteenService.GetAllCanteen(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, canteens)
	// TODO: answer here
}

//	Update canteen
//
// @Summary Update canteen
// @Description Update canteen
// @Security BearerAuth
// @Tags Canteen
// @Security Authorization
// @Accept json
// @Produce json
// @Param id path int true "Canteen id"
// @Param Canteen body model.CanteenData true "Canteen"
// @Success 200 {object} model.Canteen
// @Failure 400 {object} model.ErrorResponse
// @Router /canteens/{id} [put]
func (t *canteenAPI) Update(c *gin.Context) {
	idCanteen := c.Params.ByName("id")

	idCanteenInt, err := strconv.Atoi(idCanteen)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid Canteen id"))
		return
	}

	var Canteen model.Canteen

	if err := c.BindJSON(&Canteen); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	err = t.canteenService.UpdateCanteen(c.Request.Context(), idCanteenInt, Canteen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("internal server error"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(idCanteenInt, "Canteen updated"))
}

//	Delete canteen
//
// @Summary Delete canteen
// @Description Remove a canteen from database
// @Tags Canteen
// @Security Authorization
// @Accept json
// @Produce json
// @Param id path int true "Canteen id"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /canteens/{id} [delete]
func (t *canteenAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := t.canteenService.DeleteCanteen(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	userID, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	intUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	c.JSON(http.StatusOK, model.NewSuccessResponse(intUserID, "Canteen deleted"))
	// TODO: answer here
}
