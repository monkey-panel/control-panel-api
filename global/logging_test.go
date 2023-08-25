package global

import "testing"

func TestLog(t *testing.T) {
	Log.Debug("test")
	Log.Info("test")
	Log.Error("test")
	Log.Notice("test")
}
