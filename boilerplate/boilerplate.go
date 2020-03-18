package boilerplate

import "fmt"

var LOG_LEVEL = "error"

// Print debug message
func Debug(msg ...string) {
	if LOG_LEVEL == "debug" {
		fmt.Println(msg)
	}
}

// Check for error and panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
