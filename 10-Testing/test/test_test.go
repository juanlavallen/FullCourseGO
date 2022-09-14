package test

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(10, 20)

	if result != 30 {
		t.Error("Incorrect result")
	}
}
