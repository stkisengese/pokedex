package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name: "trailing spaces", input: "  hello  world ", want: []string{"hello", "world"},
		},
		{
			name: "caps", input: "hello WORLD", want: []string{"hello", "world"},
		},
		{
			name: "empty", input: "  ", want: []string{},
		},
		{
			name: "one word", input: " hello   ", want: []string{"hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanInput(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
