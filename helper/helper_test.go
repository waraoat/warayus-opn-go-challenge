package helper_test

import (
	"fmt"
	"io"
	"os"
	"tamboon/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		slice          []string
		val            string
		expectedOutput bool
	}{
		{[]string{"apple", "banana", "orange"}, "apple", true},
		{[]string{"apple", "banana", "orange"}, "banana", true},
		{[]string{"apple", "banana", "orange"}, "orange", true},
		{[]string{"apple", "banana", "orange"}, "grape", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("slice=%v,val=%s", test.slice, test.val), func(t *testing.T) {
			output := helper.Contains(test.slice, test.val)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestShowProgressBar(t *testing.T) {
	tests := []struct {
		current int
		total   int
		expectedOutput  string
		
	}{
		{0, 10, "\r[                                                  ] (0/10)"},
		{5, 10, "\r[=========================                         ] (5/10)"},
		{10, 10, "\r[==================================================] (10/10)"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("current=%d,total=%d", test.current, test.total), func(t *testing.T) {
			output := captureOutput(func() {
				helper.ShowProgressBar(test.current, test.total)
			})

			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
