package database

import (
	"errors"
	"strings"

	"github.com/a3510377/control-panel-api/common/codes"
	. "github.com/a3510377/control-panel-api/common/types"
	"github.com/a3510377/control-panel-api/common/utils"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type UserInfo struct {
	BaseModel
	Nickname    string        `json:"nickname"`
	Permissions Permission    `json:"permissions"`
	Token       *RefreshToken `json:"token,omitempty"`
	Lang        string        `json:"lang,omitempty"`
}

// login user struct
type LoginUser struct {
	Username string `json:"username" form:"username" binding:"required,lowercase,alphanum,min=4,max=20" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password" binding:"required,min=4,max=20"`
}

// create new user struct
type NewUser struct {
	LoginUser
	Nickname string `json:"nickname" form:"nickname" binding:"required,max=32"`
}

// in struct add token
func (u *UserInfo) AttachToken() { u.Token = NewJWT(u.ID) }

// get user from name
func (d DB) GetUserFromName(name string) *DBUser {
	user := &DBUser{}
	err := d.Where("username = ?", name).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}

// get user from id
func (d DB) GetUserFromID(id ID) *DBUser {
	user := &DBUser{}
	err := d.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}

// get user from token
func (d DB) GetUserFromToken(token string) *DBUser {
	if claims := JWT(token).Info(); claims != nil {
		return d.GetUserFromID(claims.ID)
	}

	return nil
}

// create new user
func (d DB) CreateUser(user NewUser) (*UserInfo, error) {
	if err := binding.Validator.ValidateStruct(user); err != nil {
		return nil, err
	}

	user.Password = utils.BcryptHash(user.Password)
	data := DBUser{
		LoginUser: user.LoginUser,
		Nickname:  user.Nickname,
	}

	if err := d.Create(&data).Error; err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return nil, codes.UsernameAlreadyExists
		}
		return nil, err
	}

	return data.ToUserInfo(), nil
}
