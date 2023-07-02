package database

import (
	. "github.com/monkey-panel/control-panel-api/common/types"
	. "github.com/monkey-panel/control-panel-utils/types"

	"gorm.io/gorm"
)

// set up models
func setupModel(db *DB) {
	db.AutoMigrate(&DBUser{})
}

// database base model struct
type BaseModel struct {
	ID        ID   `gorm:"primarykey" json:"id"`
	CreatedAt Time `gorm:"<-:create;autoCreateTime" json:"create_at"`
}

// set ID
func (i *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = GlobalIDMake.Generate()
	return
}

// database user struct
type DBUser struct {
	BaseModel
	LoginUser
	ID          ID         `gorm:"primarykey;many2many:user_instance" json:"id"`
	Nickname    string     `json:"nickname" binding:"min=1,max=32"`
	Lang        string     `json:"lang,omitempty"`
	Permissions Permission `json:"permissions"`
}

/* set up table name */
func (DBUser) TableName() string { return "user" }

func (u DBUser) ToUserInfo() *UserInfo {
	return &UserInfo{
		BaseModel:   BaseModel{ID: u.ID, CreatedAt: u.CreatedAt},
		Nickname:    u.Nickname,
		Permissions: u.Permissions,
		Lang:        u.Lang,
	}
}
