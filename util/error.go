package util

import "log"

// ErrorHandle handle
func ErrorHandle(err error) {
	if err != nil {
		log.Println(err)
	}
}
