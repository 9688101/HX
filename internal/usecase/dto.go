package usecase

type RegisterUserRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email"`             // 可选，若开启邮箱验证则必填
	VerificationCode string `json:"verification_code"` // 验证码，邮箱验证时必填
	AffCode          string `json:"aff_code"`          // 邀请码
}
