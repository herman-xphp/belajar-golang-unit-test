package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldEko(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldAnto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anto")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Ucok", func(t *testing.T) {
		result := HelloWorld("Ucup")
		require.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
	})
	t.Run("Anto", func(t *testing.T) {
		result := HelloWorld("Anto")
		require.Equal(t, "Hello Anto", result, "Result must be 'Hello Anto'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Ucup")
	require.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
}

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
