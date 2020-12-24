package logger

import "testing"

func TestMustSetup(t *testing.T) {
	MustSetup(LogConf{
		ServiceName: "api",
		Colorful:    true,
	})
	for i := 0; i < 10; i++ {
		Infof("123")
	}
}
