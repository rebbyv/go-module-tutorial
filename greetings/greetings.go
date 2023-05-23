package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var timeNow string

func init() {
	rand.Seed(time.Now().UnixNano())
	timeNow = time.Now().Format("15:04")
}

func Hello (name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomGreeting(), name, timeNow)
	return message, nil
}

func randomGreeting () string {
	formats := []string{
		"Hi, %v. Welcome! It's currently %v.",
		"Great to see you, %v! It's currently %v.",
		"Hail, %v! Well met! It's currently %v.",
	}
	return formats[rand.Intn(len(formats))]
}