package main

import "testing"

func TestHelloWorld(t *testing.T) {
	whatWeGot := HelloWorld()
	whatWeExpect := "Hello World!"
	if whatWeGot != whatWeExpect {
		t.Error("got:", whatWeGot, "expected:", whatWeExpect)
	}
}
