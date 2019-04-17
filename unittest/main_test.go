package main

import "testing"

// pass
func TestDouble(t *testing.T) {
	var testNum = 4
	result := Double(testNum)
	if result != testNum*2 {
		t.Fail()
	}
}

// fail
func TestTriple(t *testing.T) {
	var testNum = 4
	result := Triple(testNum)
	if result != testNum*3 {
		t.Fail()
	}
}
