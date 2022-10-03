package device

import (
	"fmt"
	"testing"
)

func TestGetDevices(t *testing.T) {
	GetDevices()

}
func TestRandomNumb(t *testing.T) {

	GetRandomDevice()
}
func TestGetRandomDevice(t *testing.T) {
	GetDevices()
	fmt.Println(GetRandomDevice())

}
