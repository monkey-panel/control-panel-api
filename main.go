package main

import (
	"math/rand"
	"time"

	"github.com/a3510377/control-panel-api/database"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	database.NewDB("db.db")
}
