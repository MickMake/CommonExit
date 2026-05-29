package main

import (
	"errors"
	"fmt"

	"github.com/MickMake/CommonExit"
)

func main() {
	value, err := buildValue("deck")
	fmt.Printf("value=%q err=%v\n", value, err)
}

func buildValue(input string) (value string, err error) {
	for range CommonExit.Block {
		if input == "" {
			err = errors.New("empty input")
			break
		}

		value = "job:" + input
	}

	return value, err
}
