package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	var tests = []struct {
		a          bytes.Buffer
		countLines bool
		countBytes bool
		want       int
	}{
		{*bytes.NewBufferString("word1 word2 word3 word4\n"), false, false, 4},
		{*bytes.NewBufferString("word1\n"), false, false, 1},
		{*bytes.NewBufferString(""), false, false, 0},
		{*bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1"), true, false, 3},
		{*bytes.NewBufferString("gopher"), false, true, 6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Count(&tt.a, tt.countLines, tt.countBytes)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
