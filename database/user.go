package database

// login user struct
type LoginUser struct {
	Username string `json:"username" validate:"required,min=4,max=20" gorm:"uniqueIndex"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

// create new user struct
type NewUser struct {
	LoginUser
}
