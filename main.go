package fgp

import "fmt"

// Println is a simple wrapper around fmt.Println() that adds a prefix
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println("[myprint] ", a)
}
