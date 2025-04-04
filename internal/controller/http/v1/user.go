package v1

// **处理 /api/currentUser**
// func (ctrl *UserController) GetCurrentUserHandler(c *gin.Context) {
// 	// 伪造的用户数据库（实际应用中应该查询数据库）
// 	var mockUsers = map[string]entity.User{
// 		"user_token_123": {
// 			Name:      "张三",
// 			Avatar:    "https://example.com/avatar.png",
// 			UserID:    "user_001",
// 			Email:     "zhangsan@example.com",
// 			Signature: "坚持就是胜利",
// 			Title:     "普通用户",
// 			Group:     "user",
// 			// Tags:        []Tag{{Key: "role", Label: "普通用户"}}
// 			NotifyCount: 2,
// 			UnreadCount: 5,
// 			Country:     "中国",
// 			Access:      "user",
// 			// Geographic: &Geographic{
// 			// 	Province: &Location{Label: "广东省", Key: "gd"},
// 			// 	City:     &Location{Label: "深圳市", Key: "sz"},
// 			// },
// 			Address: "深圳市南山区科技园",
// 			Phone:   "13800000000",
// 		},
// 	}
// 	// 获取请求头中的 Authorization
// 	authHeader := c.GetHeader("Authorization")
// 	if authHeader == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
// 		return
// 	}

// 	// 解析 Token
// 	token := strings.TrimPrefix(authHeader, "Bearer ")

// 	// 根据 Token 查询用户
// 	user, exists := mockUsers[token]
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
// 		return
// 	}

// 	// 返回用户数据
// 	c.JSON(http.StatusOK, user)
// }

// // LoginHandler 处理用户密码登录请求
// func (ctrl *UserController) LoginHandler(c *gin.Context) {
// 	if !dyncfg.PasswordLoginEnabled {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Message: "管理员关闭了密码登录",
// 			Success: false,
// 		})
// 		return
// 	}

// 	var req usecase.LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "error",
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	// 校验参数非空（binding 已确保必填，但可以额外检查）
// 	if req.Username == "" || req.Password == "" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  "error",
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	// 调用 UseCase 层处理登录逻辑
// 	user, err := ctrl.usecase.Login(c.Request.Context(), req.Username, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  "error",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// 设置登录会话
// 	ctrl.setupLogin(user, c)
// }

// RegisterUserHandler 处理用户注册请求
// func (ctrl *UserController) RegisterUserHandler(c *gin.Context) {
// 	// 检查系统配置是否允许注册
// 	if !dyncfg.RegisterEnabled {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Success: false,
// 			Message: "管理员关闭了新用户注册",
// 		})
// 		return
// 	}
// 	if !dyncfg.PasswordRegisterEnabled {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Success: false,
// 			Message: "管理员关闭了通过密码进行注册，请使用第三方账户验证的形式进行注册",
// 		})
// 		return
// 	}

// 	// 解析 JSON 请求体到 DTO
// 	var req usecase.RegisterUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, BaseResponse{
// 			Success: false,
// 			Message: i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	// 进行输入验证
// 	if err := valid.ValidateStruct(&req); err != nil {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Success: false,
// 			Message: i18n.Translate(c, "invalid_input"),
// 		})
// 		return
// 	}

// 	// 调用 UseCase 层处理注册逻辑
// 	if err := ctrl.usecase.RegisterUser(c.Request.Context(), req); err != nil {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Success: false,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, BaseResponse{
// 		Success: true,
// 		Message: "",
// 	})
// }

// // setupLogin 设置 session，并返回清理后的用户信息
// func (ctrl *UserController) setupLogin(user *entity.User, c *gin.Context) {
// 	session := sessions.Default(c)
// 	session.Set("id", user.Id)
// 	session.Set("username", user.Username)
// 	session.Set("role", user.Role)
// 	session.Set("status", user.Status)
// 	if err := session.Save(); err != nil {
// 		c.JSON(http.StatusOK, BaseResponse{
// 			Message: "无法保存会话信息，请重试",
// 			Success: false,
// 		})
// 		return
// 	}

// 	// 返回清理后的用户信息，去除敏感字段
// 	// cleanUser := struct {
// 	// 	Id          int    `json:"id"`
// 	// 	Username    string `json:"username"`
// 	// 	DisplayName string `json:"display_name"`
// 	// 	Role        int    `json:"role"`
// 	// 	Status      int    `json:"status"`
// 	// }{
// 	// 	Id:          user.Id,
// 	// 	Username:    user.Username,
// 	// 	DisplayName: user.DisplayName,
// 	// 	Role:        user.Role,
// 	// 	Status:      user.Status,
// 	// }
// 	c.JSON(http.StatusOK, LoginResponse{
// 		Status:           "ok",
// 		Type:             "account",
// 		CurrentAuthority: "user",
// 	})
// }

// // // LogoutHandler 处理用户注销请求，清除 session
// func (ctrl *UserController) LogoutHandler(c *gin.Context) {
// 	fmt.Println("LogoutHandler")
// 	session := sessions.Default(c)
// 	session.Clear()
// 	if err := session.Save(); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": err.Error(),
// 			"success": false,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "",
// 		"success": true,
// 	})
// }

// // // GetAllUsersHandler 处理获取用户列表请求
// func (ctrl *UserController) GetUserListHandler(c *gin.Context) {
// 	p, _ := strconv.Atoi(c.Query("p"))
// 	if p < 0 {
// 		p = 0
// 	}
// 	order := c.DefaultQuery("order", "")
// 	offset := 10
// 	limit := 10

// 	users, err := ctrl.usecase.GetUserList(c.Request.Context(), offset, limit, order)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    users,
// 	})
// }

// // // SearchUsersHandler 处理用户搜索请求
// func (ctrl *UserController) SearchUsersHandler(c *gin.Context) {
// 	keyword := c.Query("keyword")
// 	users, err := ctrl.usecase.SearchUsers(c.Request.Context(), keyword)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    users,
// 	})
// }

// // // GetUserHandler 处理根据 ID 获取单个用户信息的请求
// func (ctrl *UserController) GetUserHandler(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	callerRole := c.GetInt(consts.Role)
// 	user, err := ctrl.usecase.GetUser(c.Request.Context(), id, callerRole)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    user,
// 	})
// }

// // // UpdateSelfHandler 处理当前用户更新自己信息的请求
// func (ctrl *UserController) UpdateSelfHandler(c *gin.Context) {
// 	var req usecase.UpdateSelfRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	// 获取当前用户 ID，从上下文中（例如 ctxkey.Id）
// 	userID := c.GetInt(consts.Id)
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}

// 	if err := ctrl.usecase.UpdateSelf(c.Request.Context(), req, userID); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 	})
// }

// // // GetSelfHandler 处理获取当前用户信息的请求
// func (ctrl *UserController) GetSelfHandler(c *gin.Context) {
// 	userID := c.GetInt(consts.Id)
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}

// 	user, err := ctrl.usecase.GetSelf(c.Request.Context(), userID)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    user,
// 	})
// }

// // // DeleteSelfHandler 处理当前用户自删除的请求
// func (ctrl *UserController) DeleteSelfHandler(c *gin.Context) {
// 	userID := c.GetInt("id")
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}
// 	if err := ctrl.usecase.DeleteSelf(c.Request.Context(), userID); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "账户已删除",
// 	})
// }

// // UpdateUserHandler 处理管理员更新用户的请求
// func (ctrl *UserController) UpdateUserHandler(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	var req usecase.UpdateUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	callerRole := c.GetInt(consts.Role)
// 	if callerRole == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取调用者权限",
// 		})
// 		return
// 	}

// 	if err := ctrl.usecase.UpdateUser(ctx, req, callerRole); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 	})
// }

// // DeleteUserHandler 处理管理员删除用户的请求
// func (ctrl *UserController) DeleteUserHandler(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "无效的用户 ID",
// 		})
// 		return
// 	}

// 	callerRole := c.GetInt(consts.Role)
// 	if callerRole == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取调用者权限",
// 		})
// 		return
// 	}

// 	if err := ctrl.usecase.DeleteUser(ctx, id, callerRole); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "用户删除成功",
// 	})
// }

// // ManageUserHandler 处理管理员管理用户的请求
// func (ctrl *UserController) ManageUserHandler(c *gin.Context) {
// 	var req usecase.ManageRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	callerRole := c.GetInt("role")
// 	if callerRole == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取调用者权限",
// 		})
// 		return
// 	}

// 	updatedUser, err := ctrl.usecase.ManageUser(c.Request.Context(), req, callerRole)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	// 如果删除操作成功，updatedUser 可能为 nil，此处按更新处理
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    updatedUser,
// 	})
// }

// // EmailBindHandler 处理邮箱绑定请求
// func (ctrl *UserController) EmailBindHandler(c *gin.Context) {
// 	email := c.Query("email")
// 	code := c.Query("code")
// 	if email == "" || code == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}
// 	userID := c.GetInt("id")
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}
// 	if err := ctrl.usecase.BindEmail(c.Request.Context(), email, code, userID); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 	})
// }

// // GenerateAccessTokenHandler 处理生成访问令牌的请求
// func (ctrl *UserController) GenerateAccessTokenHandler(c *gin.Context) {
// 	userID := c.GetInt(consts.Id)
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}
// 	token, err := ctrl.usecase.GenerateAccessToken(c.Request.Context(), userID)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    token,
// 	})
// }

// // GetAffCodeHandler 处理获取邀请码的请求
// func (ctrl *UserController) GetAffCodeHandler(c *gin.Context) {
// 	userID := c.GetInt(consts.Id)
// 	if userID == 0 {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"success": false,
// 			"message": "无法获取当前用户信息",
// 		})
// 		return
// 	}
// 	affCode, err := ctrl.usecase.GetAffCode(c.Request.Context(), userID)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    affCode,
// 	})
// }

// func (ctrl *UserController) ResetPassword(c *gin.Context) {
// 	var req usecase.PasswordResetRequest
// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if req.Email == "" || req.Token == "" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}
// 	if !verif.VerifyCodeWithKey(req.Email, req.Token, verif.PasswordResetPurpose) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "重置链接非法或已过期",
// 		})
// 		return
// 	}
// 	password := verif.GenerateVerificationCode(12)
// 	// err = ctrl.usecase.ResetUserPasswordByEmail(req.Email, password)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	verif.DeleteKey(req.Email, verif.PasswordResetPurpose)
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    password,
// 	})
// 	return
// }
