package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() //mengeksekusi semua unit test

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
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
		// t.Fail() //tetap mengeksekusi kode dibawahnya
		t.Error("Result must be 'Hello Marlin'")
	}
	fmt.Println("TestHelloWorldMarlin Done")
}

func TestHelloWorldPurnama(t *testing.T) {
	result := HelloWorld("Purnama")

	if result != "Hello Purnama" {
		//error
		// t.FailNow() //jika fail, kode dibawahnya tidak lagi dieksekusi langsung berhenti
		t.Fatal("Result must be 'Hello Purnama'")
	}
	fmt.Println("TestHelloWorldPurnama Done")
}
