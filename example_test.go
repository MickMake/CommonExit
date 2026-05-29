package CommonExit_test

import (
	"errors"
	"fmt"

	"github.com/MickMake/CommonExit"
)

func ExampleBlock_success() {
	var err error
	result := ""

	for range CommonExit.Block {
		err = validateName("Mick")
		if err != nil {
			break
		}

		result = "ready"
	}

	fmt.Println(result, err)
	// Output: ready <nil>
}

func ExampleBlock_bailout() {
	var err error
	result := "not started"

	for range CommonExit.Block {
		err = validateName("")
		if err != nil {
			break
		}

		result = "ready"
	}

	fmt.Println(result, err)
	// Output: not started empty name
}

func validateName(name string) error {
	if name == "" {
		return errors.New("empty name")
	}
	return nil
}
