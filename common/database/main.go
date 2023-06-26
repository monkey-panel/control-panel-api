package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	Validate *validator.Validate
}

func NewDB(filename string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time { return time.Now().UTC() },
	})

	new_db := &DB{db, validator.New()}

	setupModel(new_db)

	return new_db, err
}

func GetDBFromContext(ctx *gin.Context) *DB {
	return ctx.MustGet("DB").(*DB)
}
