package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/guntoroyk/go-user-api/config"
	"github.com/guntoroyk/go-user-api/entity"
	httpHandler "github.com/guntoroyk/go-user-api/handler/http"
	"github.com/guntoroyk/go-user-api/lib/jwt"
	"github.com/guntoroyk/go-user-api/lib/validator"
	customMiddleware "github.com/guntoroyk/go-user-api/middleware"
	dbRepo "github.com/guntoroyk/go-user-api/repository/db"
	proxyRepo "github.com/guntoroyk/go-user-api/repository/proxy"
	"github.com/guntoroyk/go-user-api/storage"
	"github.com/guntoroyk/go-user-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	e := echo.New()
	cMiddleware := customMiddleware.NewMiddleware()

	validator := validator.GetValidator()
	db := storage.NewDB(&config.DB)
	tokenService := jwt.NewTokenService(config.App.Secret)

	// Repositories
	userRepo := dbRepo.NewUserRepo(db)
	userRepoProxy := proxyRepo.NewUserRepoProxy(userRepo)

	// Usecases
	userUsecase := usecase.NewUserUsecase(userRepoProxy, validator)
	authUsecase := usecase.NewAuthUsecase(userRepo, tokenService, validator)

	// Handlers
	handler := httpHandler.NewHandler(userUsecase, authUsecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/login", handler.Login)

	e.File("/docs", "public/go-user-api-postman_collection.json")

	jwtConfig := middleware.JWTConfig{
		Claims:     &entity.JwtCustomClaims{},
		SigningKey: []byte(config.App.Secret),
	}
	auth := e.Group("", middleware.JWTWithConfig(jwtConfig))

	auth.GET("/users", handler.GetUsers, cMiddleware.Authorize(entity.RoleAdmin, entity.RoleUser))
	auth.GET("/users/:id", handler.GetUser, cMiddleware.Authorize(entity.RoleAdmin, entity.RoleUser))

	auth.POST("/users", handler.CreateUser, cMiddleware.Authorize(entity.RoleAdmin))
	auth.PATCH("/users/:id", handler.UpdateUser, cMiddleware.Authorize(entity.RoleAdmin))
	auth.DELETE("/users/:id", handler.DeleteUser, cMiddleware.Authorize(entity.RoleAdmin))

	e.Logger.Fatal(e.Start(config.App.Host + ":" + config.App.Port))
}
