package api

import (
	"net/http"

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserAPI interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

// Register is the handler for the register endpoint.
// @Summary Perform user register
// @Description Register a user and returns an access token
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body model.UserCredential true "User credentials"
// @Success 200 {object} model.UserAuthResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/register [post]
func (u *userAPI) Register(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Username: user.Username,
		Password: user.Password,
	}

	// check if username already exist
	registeredUser, _ := u.userService.GetByUsername(c.Request.Context(), user.Username)
	if registeredUser.Username != "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("username already exist"))
		return
	}

	recordUser, err := u.userService.Register(c.Request.Context(), &recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	id, _ := u.userService.Login(c.Request.Context(), &recordUser)

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     10000000000000,
	})
	token, err := jwt.SignedString([]byte("secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewAccessTokenResponse(recordUser.Username, token))
}

// Login is the handler for the login endpoint.
// @Summary Perform user login
// @Description Logs in a user and returns an access token
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body model.UserCredential true "User credentials"
// @Success 200 {object} model.UserAuthResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/login [post]
func (u *userAPI) Login(c *gin.Context) {
	// TODO: answer here
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	// get user id from database
	userData := model.User{
		Username: user.Username,
		Password: user.Password,
	}
	id, err := u.userService.Login(c.Request.Context(), &userData)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     10000000000000,
	})
	token, err := jwt.SignedString([]byte("secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, model.NewAccessTokenResponse(userData.Username, token))
}
