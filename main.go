package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	_ "a21hc3NpZ25tZW50/docs"
	api "a21hc3NpZ25tZW50/handler"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
	"a21hc3NpZ25tZW50/utils"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler        api.UserAPI
	CanteenAPIHandler     api.CanteenAPI
	MenuAPIHandler        api.MenuAPI
	CanteenMenuAPIHandler api.CanteenMenuAPI
}

// @title Canteen API
// @version 1.00
// @description This is a documentation of Canteen API.
// @description Token can be obtained from /api/v1/login and/or /api/v1/register.

// @Schemes https http
// @host canteen-api.up.railway.app
// @BasePath /api/v1

// @SecurityDefinitions.apiKey  Authorization
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		// temporary handle unsupported env files in railway.app
		log.Println("Error loading .env file, will use default env variables")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()

		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())

		err := utils.ConnectDB()
		if err != nil {
			panic(err)
		}

		db := utils.GetDBConnection()

		router = RunServer(db, router)

		fmt.Println("Server is running on port 8080")
		err = router.Run()
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	userRepo := repository.NewUserRepository(db)
	canteenRepo := repository.NewCanteenRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	canteenMenuRepo := repository.NewCanteenMenuRepository(db)

	userService := service.NewUserService(userRepo)
	canteenService := service.NewCanteenService(canteenRepo, userRepo)
	menuService := service.NewMenuService(menuRepo)
	canteenMenuService := service.NewCanteenMenuService(canteenMenuRepo)

	userAPIHandler := api.NewUserAPI(userService)
	canteenAPIHandler := api.NewCanteenAPI(canteenService)
	menuAPIHandler := api.NewMenuAPI(menuService)
	canteenMenuAPIHandler := api.NewCanteenMenuAPI(canteenMenuService)

	apiHandler := APIHandler{
		UserAPIHandler:        userAPIHandler,
		CanteenAPIHandler:     canteenAPIHandler,
		MenuAPIHandler:        menuAPIHandler,
		CanteenMenuAPIHandler: canteenMenuAPIHandler,
	}

	gin.Use(middleware.Cors())

	users := gin.Group("/api/v1/users")
	{
		users.POST("/login", apiHandler.UserAPIHandler.Login)
		users.POST("/register", apiHandler.UserAPIHandler.Register)
	}

	canteens := gin.Group("/api/v1/canteens")
	{
		canteens.GET("/", apiHandler.CanteenAPIHandler.Get)

		canteens.Use(middleware.Auth())
		canteens.POST("/", apiHandler.CanteenAPIHandler.Create)
		canteens.GET("seed/:location", apiHandler.CanteenAPIHandler.Seed)
		canteens.PUT("/:id", apiHandler.CanteenAPIHandler.Update)
		canteens.DELETE("/:id", apiHandler.CanteenAPIHandler.Delete)
	}

	menus := gin.Group("/api/v1/menus")
	{
		menus.GET("/", apiHandler.MenuAPIHandler.GetAll)
		menus.GET("/:id", apiHandler.MenuAPIHandler.GetByID)

		menus.Use(middleware.Auth())
		menus.POST("/", apiHandler.MenuAPIHandler.Create)
		menus.PUT("/:id", apiHandler.MenuAPIHandler.Update)
		menus.DELETE("/:id", apiHandler.MenuAPIHandler.Delete)
	}

	canteenMenus := gin.Group("/api/v1/canteen-menus")
	{
		canteenMenus.GET("/", apiHandler.CanteenMenuAPIHandler.GetAll)
		canteenMenus.GET("/:id", apiHandler.CanteenMenuAPIHandler.GetByCanteenID)

		canteenMenus.Use(middleware.Auth())
		canteenMenus.POST("/", apiHandler.CanteenMenuAPIHandler.Create)
		canteenMenus.PUT("/:id", apiHandler.CanteenMenuAPIHandler.Update)
		canteenMenus.DELETE("/:id", apiHandler.CanteenMenuAPIHandler.Delete)
	}

	gin.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return gin
}
