package testmod

import (
	"errors"
	"fmt"
)

func HelloModule(message string, name string, hoge ...string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty.")
	}
	return fmt.Sprintf("Hello Gomodule %v !! (%v)", message, name), nil
}
