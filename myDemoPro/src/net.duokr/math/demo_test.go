package math

import (
	"fmt"
	"testing"
)

func TestAddMethod(t *testing.T) {
	fmt.Println("doing TestAddMethod()")
	var a, b = 10, 20
	var res = AddMethod(a, b)
	if res != a+b {
		t.Error("Test Case [ TestAddMethod ] Failed!")
	}
}
