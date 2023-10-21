package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CanteenMenuAPI interface {
	Create(c *gin.Context)
	GetByCanteenID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type canteenMenuAPI struct {
	canteenMenuService service.CanteenMenuService
}

func NewCanteenMenuAPI(canteenService service.CanteenMenuService) *canteenMenuAPI {
	return &canteenMenuAPI{canteenService}
}

// Create Canteen Menu
// @Summary Create Canteen Menu
// @Description Link Menu to Canteen
// @Tags Canteen Menu
// @Security Authorization
// @Accept json
// @Produce json
// @Param CanteenMenu body model.CanteenMenuData true "Canteen Menu"
// @Success 200 {object} model.CanteenMenu
// @Failure 400 {object} model.ErrorResponse
// @Router /canteen-menus [post]
func (t *canteenMenuAPI) Create(c *gin.Context) {
	canteenMenu := model.CanteenMenu{}
	if err := c.BindJSON(&canteenMenu); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	err := t.canteenMenuService.CreateCanteenMenu(c.Request.Context(), canteenMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, canteenMenu)
}

// Get All Canteen Menu
// @Summary Get All Canteen Menu
// @Description Get All list of Canteen with the Menu
// @Tags Canteen Menu
// @Produce json
// @Success 200 {object} []model.CanteenMenuDetail
// @Failure 500 {object} model.ErrorResponse
// @Router /canteen-menus [get]
func (t *canteenMenuAPI) GetAll(c *gin.Context) {
	canteenMenu, err := t.canteenMenuService.GetAllCanteenMenu(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, canteenMenu)
}

// Get Canteen Menu
// @Summary Get Canteen Menu by Cantten ID
// @Description Get Canteen Menu by Canteen ID
// @Tags Canteen Menu
// @Produce json
// @Param id path int true "Canteen ID"
// @Success 200 {object} []model.CanteenMenu
// @Failure 500 {object} model.ErrorResponse
// @Router /canteen-menus/{id} [get]
func (t *canteenMenuAPI) GetByCanteenID(c *gin.Context) {
	canteenIDParam := c.Params.ByName("id")
	canteenID, _ := strconv.Atoi(canteenIDParam)

	canteenMenu, err := t.canteenMenuService.GetMenuByCanteenID(c.Request.Context(), canteenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, canteenMenu)
}

// Update Canteen Menu
// @Summary Update Canteen Menu
// @Description Update Relation
// @Tags Canteen Menu
// @Security Authorization
// @Produce json
// @Param id path int true "CanteenMenu ID"
// @Param body body model.CanteenMenuData true "Canteen Menu"
// @Success 200 {object} model.CanteenMenu
// @Failure 400 {object} model.ErrorResponse
// @Router /canteen-menus/{id} [put]
func (t *canteenMenuAPI) Update(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.Atoi(idParam)

	canteenMenu := model.CanteenMenu{}
	if err := c.BindJSON(&canteenMenu); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	err := t.canteenMenuService.UpdateCanteenMenu(c.Request.Context(), id, canteenMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, canteenMenu)
}

// Remove Canteen Menu
// @Summary Remove Canteen Menu
// @Description Remove Relation from Menu to Canteen
// @Tags Canteen Menu
// @Security Authorization
// @Produce json
// @Param id path int true "CanteenMenu ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /canteen-menus/{id} [delete]
func (t *canteenMenuAPI) Delete(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.Atoi(idParam)

	err := t.canteenMenuService.DeleteCanteenMenu(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	userID, _ := c.Get("id")
	intUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	c.JSON(http.StatusOK, model.NewSuccessResponse(intUserID, "Data successfully deleted"))
}
