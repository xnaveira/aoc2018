package day9

import (
	"fmt"
	"testing"
)

func TestMarble(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		{"example.txt", "32"},
		{"example1.txt", "8317"},
		{"example2.txt", "146373"},
		{"example3.txt", "2764"},
		{"example4.txt", "54718"},
		{"example5.txt", "37305"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s:", tc.input), func(t *testing.T) {
			r1, _, err := Run(tc.input)
			if err != nil {
				t.Fatalf("%s", err)
			}
			if r1 != tc.output {
				t.Fatalf("expected %s got %s", tc.output, r1)
			}
		})
	}
}
