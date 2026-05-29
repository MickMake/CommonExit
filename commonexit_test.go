package CommonExit_test

import (
	"errors"
	"testing"

	"github.com/MickMake/CommonExit"
)

func TestBlockRunsOnce(t *testing.T) {
	count := 0

	for range CommonExit.Block {
		count++
	}

	if count != 1 {
		t.Fatalf("expected block to run once, got %d", count)
	}
}

func TestBlockSupportsCommonExit(t *testing.T) {
	expected := errors.New("stop here")
	var err error
	steps := 0

	for range CommonExit.Block {
		steps++

		err = expected
		if err != nil {
			break
		}

		steps++
	}

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
	if steps != 1 {
		t.Fatalf("expected one completed step before break, got %d", steps)
	}
}
