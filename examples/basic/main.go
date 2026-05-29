package main

import (
	"errors"
	"fmt"

	"github.com/MickMake/CommonExit"
)

func main() {
	if err := run("timber"); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("ok")
}

func run(item string) error {
	var err error

	for range CommonExit.Block {
		err = requireItem(item)
		if err != nil {
			break
		}

		err = processItem(item)
		if err != nil {
			break
		}
	}

	return err
}

func requireItem(item string) error {
	if item == "" {
		return errors.New("missing item")
	}
	return nil
}

func processItem(item string) error {
	fmt.Println("processing", item)
	return nil
}
