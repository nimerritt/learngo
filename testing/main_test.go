package main

import "testing"

func TestHello(t *testing.T) {
	var response = Hello()
	if response != "hi" {
		t.Error("Should say 'hi'")
	}
}
