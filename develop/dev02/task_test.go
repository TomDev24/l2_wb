package main

import "testing"

// - "a4bc2d5e" => "aaaabccddddde"
func Test1(t *testing.T) {
	str := unpack("a4bc2d5e")
	if str != "aaaabccddddde" {
		t.Fatalf("Test 1: a4bc2d5e\nExpected %s, got %s", "aaaabccddddde", str)
	}
}

// - "abcd" => "abcd"
func Test2(t *testing.T) {
	str := unpack("abcd")
	if str != "abcd" {
		t.Fatalf("Test 2: abcd\nExpected %s, got %s", "abcd", str)
	}
}

// - "45" => ""
func Test3(t *testing.T) {
	str := unpack("45")
	if str != "" {
		t.Fatalf("Test 3: 45\nExpected %s (empty string), got %s", "", str)
	}
}

// - "" => ""
func Test4(t *testing.T) {
	str := unpack("")
	if str != "" {
		t.Fatalf("Test 4: (empty string)\nExpected %s (empty string), got %s", "", str)
	}
}
