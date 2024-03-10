package log

import (
	"fmt"
	"os"
)

func Errorf(format string, v ...any) {
	fmt.Printf("\033[1m\033[31merror\033[0m: ")
	fmt.Printf(format, v...)
	fmt.Printf("\n")
	os.Exit(1)
}