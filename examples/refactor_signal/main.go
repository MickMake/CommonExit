package main

import (
	"fmt"

	"github.com/MickMake/CommonExit"
)

func main() {
	_ = configureAndRun()
}

func configureAndRun() error {
	var err error

	// This common-exit block is intentionally staged. If it grows much larger,
	// it is a strong hint that the block wants to become its own function.
	for range CommonExit.Block {
		err = loadConfig()
		if err != nil {
			break
		}

		err = prepareRuntime()
		if err != nil {
			break
		}

		err = runApp()
		if err != nil {
			break
		}
	}

	return err
}

func loadConfig() error {
	fmt.Println("load config")
	return nil
}

func prepareRuntime() error {
	fmt.Println("prepare runtime")
	return nil
}

func runApp() error {
	fmt.Println("run app")
	return nil
}
