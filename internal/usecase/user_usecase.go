package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg"
	"github.com/9688101/HX/pkg/random"
	"github.com/9688101/HX/pkg/utils"
)

// 注册
func (uc *userUseCase) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
	// 如果开启邮箱验证，则校验邮箱和验证码
	if config.GetAuthenticationConfig().EmailVerificationEnabled {
		if req.Email == "" || req.VerificationCode == "" {
			return errors.New("管理员开启了邮箱验证，请输入邮箱地址和验证码")
		}
		if !pkg.VerifyCodeWithKey(req.Email, req.VerificationCode, pkg.EmailVerificationPurpose) {
			return errors.New("验证码错误或已过期")
		}
	}

	// 根据邀请码获取邀请人ID（邀请码可能为空，返回0表示无邀请人）
	inviterId, _ := uc.repo.GetUserIdByAffCode(ctx, req.AffCode)

	// 如果 DisplayName 为空，则使用 Username
	displayName := req.Username

	// 构造待注册的用户对象（这里暂不对密码做处理，后续加密）
	newUser := entity.User{
		Username:    req.Username,
		Password:    req.Password,
		DisplayName: displayName,
		InviterId:   inviterId,
		AffCode:     random.GetRandomString(4),
		AccessToken: random.GetUUID(),
	}
	// 如果开启邮箱验证，则设置邮箱
	if config.GetAuthenticationConfig().EmailVerificationEnabled {
		newUser.Email = req.Email
	}

	// 调用 Repository 层插入用户记录
	if err := uc.repo.InsertUser(ctx, &newUser, inviterId); err != nil {
		return err
	}
	return nil
}

// 登录
func (uc *userUseCase) Login(ctx context.Context, username, password string) (*entity.User, error) {
	// 从 Repo 层查询用户记录（假设按用户名唯一查询）
	user, err := uc.repo.GetUserByUsername(ctx, username, true)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 使用 pkg 包中的方法验证密码（假设 user.Password 为加密后密码）
	if !utils.ValidatePasswordAndHash(password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != entity.UserStatusEnabled {
		return nil, errors.New("用户已被封禁")
	}

	return user, nil
}

// // 获取用户列表
func (uc *userUseCase) GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error) {
	return uc.repo.GetUserList(ctx, offset, limit, order)
}

// // 搜索用户
func (uc *userUseCase) SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error) {
	return uc.repo.SearchUsers(ctx, keyword)
}

// // 获取用户
func (uc *userUseCase) GetUser(ctx context.Context, id int, callerRole int) (*entity.User, error) {
	user, err := uc.repo.GetUserByID(ctx, id, false)
	if err != nil {
		return nil, err
	}
	// 权限校验：调用者的角色必须大于目标用户的角色，除非调用者为超级管理员
	if callerRole <= user.Role && callerRole != entity.RoleRootUser {
		return nil, errors.New("无权获取同级或更高权限等级用户的信息")
	}
	return user, nil
}

// // 更新用户
func (uc *userUseCase) UpdateSelf(ctx context.Context, req UpdateSelfRequest, userID int) error {
	// 查询原始用户信息（不包含敏感字段）
	originUser, err := uc.repo.GetUserByID(ctx, userID, false)
	if err != nil {
		return err
	}

	// 构造待更新数据：这里假设允许更新用户名、密码、DisplayName
	updatedUser := &entity.User{
		Id:          originUser.Id,
		Username:    req.Username,
		DisplayName: req.DisplayName,
	}

	// 处理密码：若密码为空或为占位值，则不更新密码
	updatePassword := false
	if req.Password != "" && req.Password != "$I_LOVE_U" {
		hashedPwd, err := utils.Password2Hash(req.Password)
		if err != nil {
			return err
		}
		updatedUser.Password = hashedPwd
		updatePassword = true
	}
	if !updatePassword {
		updatedUser.Password = "" // 表示不更新密码字段
	}

	return uc.repo.UpdateUser(ctx, updatedUser)
}

// // 获取当前用户信息
func (uc *userUseCase) GetSelf(ctx context.Context, userID int) (*entity.User, error) {
	return uc.repo.GetUserByID(ctx, userID, false)
}

// // 用户自删除操作，禁止删除超级管理员账户
func (uc *userUseCase) DeleteSelf(ctx context.Context, userID int) error {
	user, err := uc.repo.GetUserByID(ctx, userID, false)
	if err != nil {
		return err
	}
	if user.Role == entity.RoleRootUser {
		return errors.New("不能删除超级管理员账户")
	}
	return uc.repo.DeleteUserByID(ctx, userID)
}

func (uc *userUseCase) UpdateUser(ctx context.Context, req UpdateUserRequest, callerRole int) error {
	// 查询原始用户信息，不返回敏感数据
	originUser, err := uc.repo.GetUserByID(ctx, req.Id, false)
	if err != nil {
		return err
	}
	// 权限校验：调用者角色必须大于目标用户，否则拒绝更新
	if callerRole <= originUser.Role && callerRole != entity.RoleRootUser {
		return errors.New("无权更新同权限等级或更高权限等级的用户信息")
	}
	// 如果更新后的角色高于调用者权限，拒绝更新
	if req.Role > 0 && callerRole <= req.Role && callerRole != entity.RoleRootUser {
		return errors.New("无权将其他用户权限等级提升到大于等于自己的权限等级")
	}

	// 构造更新对象
	updatedUser := &entity.User{
		Id:          req.Id,
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Role:        req.Role,
	}

	// 处理密码更新：如果密码为 "$I_LOVE_U"，视为不更新密码
	if req.Password != "" && req.Password != "$I_LOVE_U" {
		hashedPwd, err := utils.Password2Hash(req.Password)
		if err != nil {
			return err
		}
		updatedUser.Password = hashedPwd
	} else {
		updatedUser.Password = ""
	}

	// 调用 Repo 层更新用户
	err = uc.repo.UpdateUser(ctx, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id int, callerRole int) error {
	originUser, err := uc.repo.GetUserByID(ctx, id, false)
	if err != nil {
		return err
	}
	if callerRole <= originUser.Role {
		return errors.New("无权删除同权限等级或更高权限等级的用户")
	}
	return uc.repo.DeleteUserByID(ctx, id)
}

// ManageUser 实现管理员管理用户逻辑
func (uc *userUseCase) ManageUser(ctx context.Context, req ManageRequest, callerRole int) (*entity.User, error) {
	// 根据用户名查询用户（不返回敏感字段）
	user, err := uc.repo.GetUserByUsername(ctx, req.Username, false)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 权限校验：调用者必须拥有大于目标用户的权限（超级管理员除外）
	if callerRole <= user.Role && callerRole != entity.RoleRootUser {
		return nil, errors.New("无权更新同权限等级或更高权限等级的用户信息")
	}

	// 根据不同的 action 执行操作
	switch req.Action {
	case "disable":
		if user.Role == entity.RoleRootUser {
			return nil, errors.New("无法禁用超级管理员用户")
		}
		user.Status = entity.UserStatusDisabled
	case "enable":
		user.Status = entity.UserStatusEnabled
	case "delete":
		if user.Role == entity.RoleRootUser {
			return nil, errors.New("无法删除超级管理员用户")
		}
		if err := uc.repo.DeleteUserByID(ctx, user.Id); err != nil {
			return nil, err
		}
		// 删除操作后直接返回 nil
		return nil, nil
	case "promote":
		if callerRole != entity.RoleRootUser {
			return nil, errors.New("普通管理员无法提升其他用户为管理员")
		}
		if user.Role >= entity.RoleAdminUser {
			return nil, errors.New("该用户已经是管理员")
		}
		user.Role = entity.RoleAdminUser
	case "demote":
		if user.Role == entity.RoleRootUser {
			return nil, errors.New("无法降级超级管理员用户")
		}
		if user.Role == entity.RoleCommonUser {
			return nil, errors.New("该用户已经是普通用户")
		}
		user.Role = entity.RoleCommonUser
	default:
		return nil, fmt.Errorf("未知操作：%s", req.Action)
	}

	// 更新用户记录
	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	// 返回更新后的部分数据（例如角色和状态）
	clearUser := &entity.User{
		Id:     user.Id,
		Role:   user.Role,
		Status: user.Status,
	}
	return clearUser, nil
}

// BindEmail 实现邮箱绑定业务逻辑
func (uc *userUseCase) BindEmail(ctx context.Context, email string, code string, userID int) error {
	// 验证验证码
	if !pkg.VerifyCodeWithKey(email, code, pkg.EmailVerificationPurpose) {
		return errors.New("验证码错误或已过期")
	}
	// 查询用户信息（不返回敏感数据）
	user, err := uc.repo.GetUserByID(ctx, userID, false)
	if err != nil {
		return err
	}
	// 更新用户邮箱
	user.Email = email
	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return err
	}
	// 如果用户为超级管理员，则更新全局配置中的 RootUserEmail
	if user.Role == entity.RoleRootUser {
		// 此处假设全局配置由 config 包管理
		// config.RootUserEmail = email
	}
	return nil
}

// GenerateAccessToken 根据用户ID生成新的访问令牌，并更新用户记录
func (uc *userUseCase) GenerateAccessToken(ctx context.Context, userID int) (string, error) {
	user, err := uc.repo.GetUserByID(ctx, userID, true)
	if err != nil {
		return "", err
	}
	// 生成新令牌
	newToken := random.GetUUID()

	// 检查是否重复（此处简单示例，实际情况建议在数据库中设置唯一索引）
	// 如果数据库操作中发现重复，通常会返回错误
	user.AccessToken = newToken

	// 更新用户记录
	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return "", err
	}

	return newToken, nil
}

// GetAffCode 根据用户ID返回邀请码；若为空则生成新邀请码，并更新用户记录
func (uc *userUseCase) GetAffCode(ctx context.Context, userID int) (string, error) {
	user, err := uc.repo.GetUserByID(ctx, userID, true)
	if err != nil {
		return "", err
	}
	if user.AffCode == "" {
		user.AffCode = random.GetRandomString(4)
		if err := uc.repo.UpdateUser(ctx, user); err != nil {
			return "", err
		}
	}
	return user.AffCode, nil
}
