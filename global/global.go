package global

import (
	"github.com/gookit/slog"
	"github.com/monkey-panel/control-panel-api/common/database"
)

const (
	LogKey = "LOG"
	DBKey  = "DB"
)

var (
	Log *slog.Logger
	DB  *database.DB
)
