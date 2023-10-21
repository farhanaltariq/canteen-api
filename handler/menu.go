package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuAPI interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type menuAPI struct {
	menuService service.MenuService
}

func NewMenuAPI(menuService service.MenuService) *menuAPI {
	return &menuAPI{menuService}
}

// Create Menu
// @Summary Create Menu
// @Description Manually insert menu to database
// @Tags Menu
// @Accept json
// @Produce json
// @Param Menu body model.MenuData true "Menu"
// @Security Authorization
// @Success 200 {object} model.MenuData
// @Failure 400 {object} model.ErrorResponse
// @Router /menus [post]
func (t *menuAPI) Create(c *gin.Context) {
	menu := model.MenuData{}
	err := c.BindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	newMenu := model.Menu{
		Name:      menu.Name,
		Price:     menu.Price,
		Stock:     menu.Stock,
		ImagePath: menu.ImagePath,
	}
	err = t.menuService.CreateMenu(c.Request.Context(), newMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, menu)
}

// Get All Menu
// @Summary Get All Menu
// @Description Get All Menu from database
// @Tags Menu
// @Produce json
// @Success 200 {object} []model.Menu
// @Failure 500 {object} model.ErrorResponse
// @Router /menus [get]
func (t *menuAPI) GetAll(c *gin.Context) {
	menus, err := t.menuService.GetAllMenu(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, menus)
}

// Get Menu
// @Summary Get Menu
// @Description Menu Details
// @Tags Menu
// @Produce json
// @Param id path int true "Menu id"
// @Success 200 {object} []model.Menu
// @Failure 500 {object} model.ErrorResponse
// @Router /menus/{id} [get]
func (t *menuAPI) GetByID(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.Atoi(idParam)
	menu, err := t.menuService.GetMenuByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
	}
	if menu.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Menu with id %d not found", id)})
		return
	}
	c.JSON(http.StatusOK, menu)

}

// Update Menu
// @Summary Update Menu
// @Descriptions Update Menu By ID
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path int true "Menu id"
// @Param Menu body model.MenuData true "Menu"
// @Security Authorization
// @Success 200 {object} model.MenuData
// @Failure 400 {object} model.ErrorResponse
// @Router /menus/{id} [put]
func (t *menuAPI) Update(c *gin.Context) {
	menu := model.MenuData{}
	err := c.BindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
	}

	updatedMenu := model.Menu{
		Name:      menu.Name,
		Price:     menu.Price,
		Stock:     menu.Stock,
		ImagePath: menu.ImagePath,
	}

	idParams := c.Params.ByName("id")
	id, _ := strconv.Atoi(idParams)
	err = t.menuService.UpdateMenu(c.Request.Context(), id, updatedMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, updatedMenu)
}

// Delete Menu
// @Summary Delete Menu
// @Descriptions Delete Menu By ID
// @Tags Menu
// @Produce json
// @Param id path int true "Menu id"
// @Security Authorization
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /menus/{id} [delete]
func (t *menuAPI) Delete(c *gin.Context) {
	idParams := c.Params.ByName("id")
	id, _ := strconv.Atoi(idParams)
	err := t.menuService.DeleteMenu(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	userIDParams, _ := c.Get("id")
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", userIDParams))
	c.JSON(http.StatusOK, model.NewSuccessResponse(userID, "Data successfully deleted"))
}
