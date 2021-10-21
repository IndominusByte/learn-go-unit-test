package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloNameTable(t *testing.B) {
	tests := [...]struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Oman",
			request: "Oman",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(b *testing.B) {
			HelloName(test.request)
		})
	}
}

func BenchmarkHelloNameSub(t *testing.B) {
	t.Run("Oman", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			HelloName("Oman")
		}
	})

	t.Run("Pradipta", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			HelloName("Pradipta")
		}
	})
}

func BenchmarkHelloName(t *testing.B) {
	for i := 0; i < t.N; i++ {
		HelloName("oman")
	}
}

func TestTableHelloName(t *testing.T) {
	tests := [...]struct {
		name     string
		request  string
		expected string
	}{
		{"oman", "oman", "Hello oman"},
		{"pradipta", "pradipta", "Hello pradipta"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloName(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("oman", func(t *testing.T) {
		result := HelloName("oman")
		assert.Equal(t, "Hello oman", result)
	})

	t.Run("pradipta", func(t *testing.T) {
		result := HelloName("pradipta")
		assert.Equal(t, "Hello pradipta", result)
	})
}

func TestMain(t *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	t.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("cannot run on linux")
	}
	result := HelloName("oman")
	require.Equal(t, "Hello oman", result)
	fmt.Println("TestHelloNameRequire Done")
}

func TestHelloNameRequire(t *testing.T) {
	result := HelloName("oman")
	require.Equal(t, "Hello oman", result)
	fmt.Println("TestHelloNameRequire Done")
}

func TestHelloNameAssert(t *testing.T) {
	result := HelloName("oman")
	assert.Equal(t, "Hello oman", result)
	fmt.Println("TestHelloNameAssert Done")
}

func TestHelloNameOman(t *testing.T) {
	result := HelloName("oman")
	if result != "Hello oman" {
		// t.Fail()
		t.Error("result must be Hello oman")
	}

	fmt.Println("TestHelloNameOman Done")
}

func TestHelloNamePradipta(t *testing.T) {
	result := HelloName("Pradipta")
	if result != "Hello Pradipta" {
		// t.FailNow()
		t.Fatal("result must be Hello Pradipta")
	}
	fmt.Println("TestHelloNamePradipta Done")
}
