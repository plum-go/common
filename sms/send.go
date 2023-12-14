package sms

import (
	"fmt"
	"math/rand"
)

func Send(phone string) string {

	format := randomFormat()
	message := fmt.Sprintf(format, phone)
	return message
}

func SendMultiple(phones []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, phone := range phones {
		format := randomFormat()
		messages[phone] = fmt.Sprintf(format, phone)
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
