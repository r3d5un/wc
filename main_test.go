package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	var tests = []struct {
		a    bytes.Buffer
		want int
	}{
		{*bytes.NewBufferString("word1 word2 word3 word4\n"), 4},
		{*bytes.NewBufferString("word1\n"), 1},
		{*bytes.NewBufferString(""), 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Count(&tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
