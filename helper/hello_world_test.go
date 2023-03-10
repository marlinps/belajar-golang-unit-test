package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Marlin",
			request: "Marlin",
		},
		{
			name:    "Purnama",
			request: "Purnama",
		},
		{
			name:    "Marlin Purnama Sari",
			request: "Marlin Purnama Sari",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Marlin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Marlin")
		}
	})

	b.Run("Purnama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Purnama")
		}
	})

	//go test -v -run=NotMathUnitTest -bench=BenchmarkSub u/ menjalankan BenchmarkSub
	//go test -v -bench=BenchmarkSub/Marlin u/ menjalankan hanya sub benchmark tertentu
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Marlin")
	}
}

func BenchmarkHelloWorldPurnama(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Purnama")
	}
	//go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorldPurnama
}

func TestTableHelloWorld(t *testing.T) { //Table Test, lebih efisien dibandingkan metode testing-testing sebelumnya
	tests := []struct { //slice struct
		name     string
		request  string
		expected string
	}{
		{
			name:     "Marlin",
			request:  "Marlin",
			expected: "Hello Marlin",
		},
		{
			name:     "Purnama",
			request:  "Purnama",
			expected: "Hello Purnama",
		},
		{
			name:     "Sari",
			request:  "Sari",
			expected: "Hello Sari",
		},
		{
			name:     "Ririn",
			request:  "Ririn",
			expected: "Hello Ririn",
		},
		{
			name:     "Rayyan",
			request:  "Rayyan",
			expected: "Hello Rayyan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) { //SubTest
	t.Run("Marlin", func(t *testing.T) {
		result := HelloWorld("Marlin")
		require.Equal(t, "Hello Marlin", result, "Result must be 'Hello Marlin'")
	})
	t.Run("Purnama", func(t *testing.T) {
		result := HelloWorld("Purnama")
		require.Equal(t, "Hello Purnama", result, "Result must be 'Hello Purnama'")
	})

	//untuk menjalankan test di subtest tertentu perintah contoh go test -run TestSubTest/Marlin
}

func TestMain(m *testing.M) { //Before and After Test
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() //mengeksekusi semua unit test

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) { //TestSkip
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS ")
	}
	result := HelloWorld("Marlin")
	require.Equal(t, "Hello Marlin", result, "Result must be Hello Marlin")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Marlin")
	require.Equal(t, "Hello Marlin", result, "Result must be Hello Marlin") //require jika gagal memanggil FailNow()
	fmt.Println("Hello World with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Marlin")
	assert.Equal(t, "Hello Marlin", result, "Result must be Hello Marlin") //assert jika gagal memanggil Fail()
	fmt.Println("Hello World with Assert Done")
}
func TestHelloWorldMarlin(t *testing.T) {
	result := HelloWorld("Marlin")

	if result != "Hello Marlin" {
		//error
		// t.Fail() //jika gagal, kode dibawahnya tetap mengeksekusi kode dibawahnya
		t.Error("Result must be 'Hello Marlin'")
	}
	fmt.Println("TestHelloWorldMarlin Done")
}

func TestHelloWorldPurnama(t *testing.T) {
	result := HelloWorld("Purnama")

	if result != "Hello Purnama" {
		//error
		// t.FailNow() //jika gagal maka kode dibawahnya tidak lagi dieksekusi langsung berhenti
		t.Fatal("Result must be 'Hello Purnama'")
	}
	fmt.Println("TestHelloWorldPurnama Done")
}
