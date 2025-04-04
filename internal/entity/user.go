package entity

import "time"

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
	RoleRootUser   = 100
)

const (
	UserStatusEnabled  = 1 // don't use 0, 0 is the default value!
	UserStatusDisabled = 2 // also don't use 0
	UserStatusDeleted  = 3
)

// User if you add sensitive fields, don't forget to clean them in setupLogin function.
// Otherwise, the sensitive information will be saved on local storage in plain text!
// type User struct {
// 	Id          int    `json:"id"`
// 	Username    string `json:"username" gorm:"unique;index" validate:"max=12"`
// 	Password    string `json:"password" gorm:"not null;" validate:"min=8,max=20"`
// 	DisplayName string `json:"display_name" gorm:"index" validate:"max=20"`

// 	Email            string `json:"email" gorm:"index" validate:"max=50"`
// 	GitHubId         string `json:"github_id" gorm:"column:github_id;index"`
// 	WeChatId         string `json:"wechat_id" gorm:"column:wechat_id;index"`
// 	LarkId           string `json:"lark_id" gorm:"column:lark_id;index"`
// 	OidcId           string `json:"oidc_id" gorm:"column:oidc_id;index"`
// 	VerificationCode string `json:"verification_code" gorm:"-:all"`                                    // this field is only for Email verification, don't save it to database!
// 	AccessToken      string `json:"access_token" gorm:"type:char(32);column:access_token;uniqueIndex"` // this token is for system management
// 	AffCode          string `json:"aff_code" gorm:"type:varchar(32);column:aff_code;uniqueIndex"`
// 	InviterId        int    `json:"inviter_id" gorm:"type:int;column:inviter_id;index"`
// 	Name             string `json:"name,omitempty"`
// 	Avatar           string `json:"avatar,omitempty"`
// 	UserID           string `json:"userid,omitempty"`
// 	Signature        string `json:"signature,omitempty"`
// 	Title            string `json:"title,omitempty"`
// 	Group            string `json:"group,omitempty"`
// 	// Tags             []Tag       `json:"tags,omitempty"`
// 	NotifyCount int    `json:"notifyCount,omitempty"`
// 	UnreadCount int    `json:"unreadCount,omitempty"`
// 	Country     string `json:"country,omitempty"`
// 	Access      string `json:"access,omitempty"`
// 	// Geographic       *Geographic `json:"geographic,omitempty"`
// 	Address string `json:"address,omitempty"`
// 	Phone   string `json:"phone,omitempty"`
// }

type User struct {
	ID        uint   `json:"userid" gorm:"primaryKey"`
	Username  string `json:"name" gorm:"unique;index"`
	Password  string `json:"password" gorm:"not null;" validate:"min=8,max=20"`
	Avatar    string `json:"avatar" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
	Phone     string `json:"phone" gorm:"type:varchar(255)"`
	Address   string `json:"address" gorm:"type:varchar(255)"`
	Group     string `json:"group" gorm:"type:varchar(255)"`
	Access    string `json:"access" gorm:"type:varchar(255)"`
	CreatedAt string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updatedAt" gorm:"autoUpdateTime"`
	Role      int    `json:"role" gorm:"type:int;default:1"`   // admin, util
	Status    int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
}

// Rule 规则列表项
type Rule struct {
	Key       int       `json:"key"`
	Disabled  bool      `json:"disabled"`
	Href      string    `json:"href"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Desc      string    `json:"desc"`
	CallNo    int       `json:"callNo"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Progress  int       `json:"progress"`
}
