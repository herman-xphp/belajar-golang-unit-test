package helper

import "testing"

func TestHelloWorldUcok(t *testing.T) {
	result := HelloWorld("Ucok")

	if result != "Hello Ucok" {
		// error
		panic("Result is not 'Hello Ucok'")
	}
}
func TestHelloWorldUcup(t *testing.T) {
	result := HelloWorld("Ucup")

	if result != "Hello Ucup" {
		// error
		panic("Result is not 'Hello Ucup'")
	}
}
