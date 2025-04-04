package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := map[string]map[string]int{
		"hello world":       {"hello": 1, "world": 1},
		"foo bar foo":       {"foo": 2, "bar": 1},
		"one two two three": {"one": 1, "two": 2, "three": 1},
	}

	for input, expected := range tests {
		result := WordCount(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("WordCount(%q) = %v; want %v", input, result, expected)
		}
	}
}

func TestIsOdd(t *testing.T) {
	tests := map[int]bool{
		1:  true,
		2:  false,
		3:  true,
		10: false,
		15: true,
	}

	for input, expected := range tests {
		result := IsOdd(input)
		if result != expected {
			t.Errorf("IsOdd(%d) = %v; want %v", input, result, expected)
		}
	}
}
