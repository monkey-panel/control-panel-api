package logging

import (
	"fmt"
	"testing"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

func TestLog(t *testing.T) {
	h1, err := handler.NewEmptyConfig(
		handler.WithLogfile("./base-info.log"),        // set log file path
		handler.WithRotateMode(rotatefile.ModeRename), // rename old log file
		handler.WithBuffSize(1),                       // 4M 4*1024*1024
		handler.WithCompress(true),                    // compression old log files
		handler.WithBackupNum(5),                      // set old log files length
	).CreateHandler()
	if err != nil {
		fmt.Printf("create slog handler error: %s\n", err.Error())
		return
	}

	f := slog.AsTextFormatter(h1.Formatter())
	f.SetTemplate("[{{datetime}}] {{level}} {{caller}} {{message}}\n")
	logs := slog.NewWithHandlers(h1)
	logs.Debug("test")
	logs.Info("test")
	logs.Error("test")
	logs.Notice("test")
}
