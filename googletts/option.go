package googletts

import (
	"errors"
	"strconv"
)

// Option contains optional parameters for Google TTS API.
type Option struct {
	Client     string
	Encoding   string
	Lang       string
	Text       string
	TextLength int
	Token      string
	TTSSpeed   float64 // voice speed (0.0 - 1.0)
}

// Validate validates option parameters.
func (o Option) Validate() error {
	if o.Text == "" {
		return errors.New("`Option.Text` must not be empty")
	}
	return nil
}

func (o Option) getClient() string {
	if o.Client != "" {
		return o.Client
	}
	return "t"
}

func (o Option) getEncoding() string {
	if o.Encoding != "" {
		return o.Encoding
	}
	return "UTF-8"
}

func (o Option) getLang() string {
	if o.Lang != "" {
		return o.Lang
	}
	return "en"
}

func (o Option) getText() string {
	return o.Text
}

func (o Option) getTextLength() string {
	if o.TextLength > 0 {
		return strconv.Itoa(o.TextLength)
	}
	return strconv.Itoa(len(o.Text))
}

func (o Option) getToken() string {
	return o.Token
}

func (o Option) getTTSSpeed() string {
	if o.TTSSpeed > 0 {
		return strconv.FormatFloat(o.TTSSpeed, 'f', 3, 64)
	}
	return "1"
}
