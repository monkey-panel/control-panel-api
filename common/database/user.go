package database

import (
	"strings"

	. "github.com/a3510377/control-panel-api/common/types"
)

type UserInfo struct {
	BaseModel
	Nickname    string     `json:"nickname"`
	Permissions Permission `json:"permissions"`
}

// login user struct
type LoginUser struct {
	Username string `json:"username" form:"username" validate:"required,min=4,max=20" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password" validate:"required,min=4,max=20"`
}

// create new user struct
type NewUser struct {
	LoginUser
	Nickname string `json:"nickname" form:"nickname" validate:"min=1,max=32"`
}

func (d DB) CreateUser(user NewUser) (*UserInfo, error) {
	if err := d.Validate.Struct(user); err != nil {
		return nil, err
	}

	data := DBUser{
		LoginUser: user.LoginUser,
		Nickname:  user.Nickname,
	}

	if err := d.Create(&data).Error; err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return nil, err
		}
	}

	return &UserInfo{
		BaseModel:   data.BaseModel,
		Nickname:    data.Nickname,
		Permissions: data.Permissions,
	}, nil
}

func (d DB) GetUserByJWT(token string) error {
	return nil
}
