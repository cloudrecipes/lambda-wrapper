// Package fs implements operations with file system.
package fs

import "io/ioutil"

// ReadFile reterns fila content or error.
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
