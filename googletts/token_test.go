package googletts

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTTSToken(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		str string
	}{
		{"a"},
		{"b"},
		{"foo"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
		{"こんにちは、世界。"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)

		token, err := GetTTSToken(tt.str)
		a.NoError(err, target)
		a.NotEmpty(token, err, target)

		parts := strings.Split(token, ".")
		a.Len(parts, 2, target)
		for _, part := range parts {
			intVal, err := strconv.Atoi(part)
			a.NoError(err, target)
			a.NotEqual(0, intVal, target)
		}
	}
}
