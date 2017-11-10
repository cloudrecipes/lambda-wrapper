// Package fs implements operations with file system.
package fs

import "io/ioutil"

// ReadFile reterns fila content or error.
func ReadFile(filename string) (string, error) {
	payload, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(payload), err
}
