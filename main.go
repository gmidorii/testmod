package testmod

import "fmt"

func HelloModule(s string) string {
	return fmt.Sprintf("Hello Gomodule %v !!", s)
}
