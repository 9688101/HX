package usecase

// 用户注册请求结构体
type RegisterUserRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email"`             // 可选，若开启邮箱验证则必填
	VerificationCode string `json:"verification_code"` // 验证码，邮箱验证时必填
	AffCode          string `json:"aff_code"`          // 邀请码
}

// 用户登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
