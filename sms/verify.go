package sms

import (
	"errors"
	"fmt"
)

func Verify(phone string, code string) (string, error) {

	if code == "" {
		return "", errors.New("code is empty")
	}
	message := fmt.Sprintf("Hi, %v. Code %v validate successful!", phone, code)
	return message, nil
}
