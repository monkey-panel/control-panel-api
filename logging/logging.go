package logging

import (
	"fmt"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

var log *slog.Logger

func init() {
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
	log = slog.NewWithHandlers(h1)
}

// Log a message with level
func Log(level slog.Level, args ...any) { log.Log(level, args) }

// Logf a format message with level
func Logf(level slog.Level, format string, args ...any) { log.Logf(level, format, args...) }

// Print logs a message at level PrintLevel
func Print(args ...any) { log.Print(args...) }

// Println logs a message at level PrintLevel
func Println(args ...any) { log.Println(args...) }

// Printf logs a message at level PrintLevel
func Printf(format string, args ...any) { log.Printf(format, args...) }

// Warn logs a message at level Warn
func Warn(args ...any) { log.Warn(args...) }

// Warnf logs a message at level Warn
func Warnf(format string, args ...any) { log.Warnf(format, args...) }

// Warning logs a message at level Warn, alias of Logger.Warn()
func Warning(args ...any) { log.Warning(args...) }

// Info logs a message at level Info
func Info(args ...any) { log.Info(args...) }

// Infof logs a message at level Info
func Infof(format string, args ...any) { log.Infof(format, args...) }

// Trace logs a message at level trace
func Trace(args ...any) { log.Trace(args) }

// Tracef logs a message at level trace
func Tracef(format string, args ...any) { log.Tracef(format, args...) }

// Error logs a message at level error
func Error(args ...any) { log.Error(args...) }

// Errorf logs a message at level error
func Errorf(format string, args ...any) { log.Errorf(format, args...) }

// ErrorT logs a error type at level error
func ErrorT(err error) { log.ErrorT(err) }

// Notice logs a message at level notice
func Notice(args ...any) { log.Notice(args...) }

// Noticef logs a message at level notice
func Noticef(format string, args ...any) { log.Noticef(format, args...) }

// Debug logs a message at level debug
func Debug(args ...any) { log.Debug(args...) }

// Debugf logs a message at level debug
func Debugf(format string, args ...any) { log.Debugf(format, args...) }

// Fatal logs a message at level fatal
func Fatal(args ...any) { log.Fatal(args...) }

// Fatalf logs a message at level fatal
func Fatalf(format string, args ...any) { log.Fatalf(format, args...) }

// Fatalln logs a message at level fatal
func Fatalln(args ...any) { log.Fatalln(args...) }

// Panic logs a message at level panic
func Panic(args ...any) { log.Panic(args...) }

// Panicf logs a message at level panic
func Panicf(format string, args ...any) { log.Panicf(format, args...) }

// Panicln logs a message at level panic
func Panicln(args ...any) { log.Panicln(args...) }
