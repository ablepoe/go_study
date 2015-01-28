package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 100
	var b = 200

	var val = Add(a, b)
	if val != a+b {
		t.Error("Test Case [", "TestAdd", "] Failed!")
	} else {
		t.Log("Test Case [", "TestAdd", "] succed!")
	}
}

func TestSubtract(t *testing.T) {
	var a = 100
	var b = 200

	var val = Subtract(a, b)
	if val != a-b {
		t.Error("Test Case [", "TestSubtract", "] Failed!")
	} else {
		t.Log("Test Case [", "TestSubtract", "] succed!")
	}

}

func TestMultiply(t *testing.T) {
	var a = 100
	var b = 200

	var val = Multiply(a, b)
	if val != a*b {
		t.Error("Test Case [", "TestMultiply", "] Failed!")
	} else {
		t.Log("Test Case [", "TestMultiply", "] succed!")
	}
}

func TestDivideNormal(t *testing.T) {
	var a = 100
	var b = 200

	var val = Divide(a, b)
	if val != a/b {
		t.Error("Test Case [", "TestDivideNormal", "] Failed!")
	} else {
		t.Log("Test Case [", "TestDivideNormal", "] succed!")
	}
}
