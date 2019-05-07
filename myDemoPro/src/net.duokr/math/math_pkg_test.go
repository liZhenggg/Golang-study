package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("TestAdd...")
	var a = 100
	var b = 200
	var val = Add(a, b)
	if val != a+b {
		t.Error("Test Case [", "TestAdd", "] Failed!")
	}
}
func TestSubtract(t *testing.T) {
	fmt.Println("TestSubtract...")
	var a = 100
	var b = 200
	var val = Subtract(a, b)
	if val != a-b {
		t.Error("Test Case [", "TestSubtract", "] Failed!")
	}
}
func TestMultiply(t *testing.T) {
	fmt.Println("TestMultiply...")
	var a = 100
	var b = 200
	var val = Multiply(a, b)
	if val != a*b {
		t.Error("Test Case [", "TestMultiply", "] Failed!")
	}
}
func TestDivideNormal(t *testing.T) {
	fmt.Println("TestDivideNormal...")
	var a = 100
	var b = 200
	var val = Divide(a, b)
	if val != a/b {
		t.Error("Test Case [", "TestDivideNormal", "] Failed!")
	}
}
