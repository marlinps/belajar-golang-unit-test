package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Marlin")

	if result != "Hello Marlin" {
		//error
		panic("Result is not 'Hello Marlin'")

	}
}

func TestHelloWorldPurnama(t *testing.T) {
	result := HelloWorld("Purnama")

	if result != "Hello Purnama" {
		//error
		panic("Result is not 'Hello Purnama'")

	}
}
