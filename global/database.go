package global

import "github.com/monkey-panel/control-panel-api/common/database"

func init() {
	db, err := database.NewDB("data/db.db")
	if err != nil {
		panic(err)
	}

	DB = db
}
