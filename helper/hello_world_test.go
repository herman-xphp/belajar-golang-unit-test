package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldUcok(t *testing.T) {
	result := HelloWorld("Ucok")

	if result != "Hello Ucok" {
		// error
		t.Error("Result must be 'Hello Ucup'")
	}

	fmt.Println("TestHelloWorldUcok")
}
func TestHelloWorldUcup(t *testing.T) {
	result := HelloWorld("Ucup")

	if result != "Hello Ucup" {
		// error
		t.Fatal("Result must be 'Hello Ucup'")
	}

	fmt.Println("TestHelloWorldUcup")
}
