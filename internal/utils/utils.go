package utils

import (
	"errors"
	"log"
)

// Log - logs errro with stack trace
func Log(err error) {
	if err != nil {
		log.Printf("%+v", errors.Unwrap(err))
	}
}
