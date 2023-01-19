package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name param can't be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formmats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
		"Yoo %v, good to see u bro!",
	}

	return formmats[rand.Intn(len(formmats))]
}
