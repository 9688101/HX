package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repo.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)

	// 注册用户接口
	r.POST("/register", userController.RegisterUserHandler)
	r.POST("/login", userController.LoginHandler)
	r.POST("/logout", userController.LogoutHandler)
	r.GET("/users", userController.GetAllUsersHandler)
	r.GET("/users/search", userController.SearchUsersHandler)
	r.GET("/user/:id", userController.GetUserHandler)

	return r
}
