package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/db"
	"github.com/9688101/HX/pkg/repo"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repo.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)

	// 注册用户注册接口
	r.POST("/register", userController.RegisterUserHandler)

	return r
}
