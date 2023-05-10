package fgp

import (
	f "fmt"
)

func init() {
	f.Println("init function - [fpg]")
}

// Println is a simple wrapper around fmt.Println() that adds a prefix
func Println(a ...interface{}) (n int, err error) {
	return f.Println("[myprint] ", a)
}
