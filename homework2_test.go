package main

import (
	"testing"
)

func TestCatchAnimal(t *testing.T) {
	fly := catchAnimal("fly")
	if len(fly) != 0 {
		t.Error("test failed")
	}
	horse := catchAnimal("horse")
	if len(horse) != 0 {
		t.Error("test failed")
	}
}
