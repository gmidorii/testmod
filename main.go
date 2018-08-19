package testmod

import (
	"errors"
	"fmt"
)

func HelloModule(s string, name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty.")
	}
	return fmt.Sprintf("Hello Gomodule %v !! (%v)", s, s), nil
}
