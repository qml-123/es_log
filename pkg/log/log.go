package log

import "github.com/qml-123/log_sdk"

func InitLogger(url []string) error {
	_, err := log_sdk.NewLogger(url[0], "log")
	if err != nil {
		return err
	}
	return nil
}
