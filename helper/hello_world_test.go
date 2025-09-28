package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ucup")
	require.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ucup")
	assert.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
	fmt.Println("TestHelloWorld with Assert Done")
}

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
