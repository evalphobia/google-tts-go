package googletts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTTSURL(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		str  string
		lang string
	}{
		{"a", "en"},
		{"b", "en"},
		{"foo", "en"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", "en"},
		{"こんにちは、世界。", "ja"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)

		u, err := GetTTSURL(tt.str, tt.lang)
		a.NoError(err, target)
		a.NotEmpty(u, err, target)
		a.Contains(u, "https://translate.google.com/translate_tts?", target)
	}
}
