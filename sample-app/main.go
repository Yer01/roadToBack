package main

import "errors"

func validateStatus(status string) error {
	size := len(status)
	if size == 0 {
		return errors.New("status cannot be empty")
	} else if size > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}
