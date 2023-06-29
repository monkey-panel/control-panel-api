package database

import (
	. "github.com/a3510377/control-panel-api/common/types"

	"gorm.io/gorm"
)

// set up models
func setupModel(db *DB) {
	db.AutoMigrate(&DBUser{}, &DBInstance{}, &DBUserInstance{})
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
	Lang        string     `json:"lang"`
	Permissions Permission `json:"permissions"`
}

// database instance struct
type DBInstance struct {
	BaseModel
	Name             string `gorm:"not null"`
	Description      string
	AdminDescription string
	AutoStart        bool
	Mark             InstanceMark
	LastAt           Time
	EndAt            Time
}

// many to many relationship between user and instance
type DBUserInstance struct {
	InstanceID ID
	UserID     ID

	Permissions Permission
	Nickname    string
}

/* set up table name */
func (DBUser) TableName() string         { return "user" }
func (DBInstance) TableName() string     { return "instance" }
func (DBUserInstance) TableName() string { return "user_instance" }

func (u DBUser) ToUserInfo() *UserInfo {
	return &UserInfo{
		BaseModel:   BaseModel{ID: u.ID, CreatedAt: u.CreatedAt},
		Nickname:    u.Nickname,
		Permissions: u.Permissions,
	}
}
