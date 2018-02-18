package googletts

import (
	"fmt"
	"net/url"
)

const (
	defaultAPIURL = "https://translate.google.com/translate_tts"
)

// GetTTSURL gets tts url for the given text.
func GetTTSURL(text, lang string) (string, error) {
	return GetTTSURLWithOption(Option{
		Text: text,
		Lang: lang,
	})
}

// GetTTSURLWithOption gets tts url with option.
func GetTTSURLWithOption(opt Option) (string, error) {
	if err := opt.Validate(); err != nil {
		return "", err
	}

	token := opt.getToken()
	if token == "" {
		var err error
		token, err = GetTTSToken(opt.Text)
		if err != nil {
			return "", err
		}
	}

	params := &url.Values{}
	params.Set("ie", opt.getEncoding())
	params.Set("client", opt.getClient())
	params.Set("tl", opt.getLang())
	params.Set("q", opt.getText())
	params.Set("textlen", opt.getTextLength())
	params.Set("ttsspeed", opt.getTTSSpeed())
	params.Set("tk", token)

	return fmt.Sprintf("%s?%s", defaultAPIURL, params.Encode()), nil
}
