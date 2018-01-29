package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if "Hello" != getHello() {
		t.Fatal("failed test")
	}
}

func TestWorld(t *testing.T) {
  if "World" != getWorld() {
    t.Fatal("failed test")
  }
}
