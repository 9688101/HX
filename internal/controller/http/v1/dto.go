package v1

import "github.com/9688101/HX/internal/usecase"

type UserController struct {
	usecase usecase.UserUseCase
}

// 创建 UserController 实例
func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		usecase: uc,
	}
}

// BaseResponse 定义了所有回复的公共字段
type BaseResponse struct {
	Success bool   `json:"success"` // 表示操作是否成功
	Message string `json:"message"` // 返回的提示信息
}

// Response 作为统一的回复结构体，可以用于返回任意复杂数据
type Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"` // 数据部分，根据需要可以是 map、结构体、数组等
}
