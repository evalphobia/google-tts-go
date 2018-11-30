package googletts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

// get token seed from HTML.
var tokenRegex = regexp.MustCompile(`,tkk:'([\d]+)\.([\d]+)',`)

// GetTTSToken gets TTS token.
func GetTTSToken(text string) (string, error) {
	if len(text) > 200 {
		return "", fmt.Errorf("text length (%d) should be less than 200 characters", len(text))
	}

	key1, key2, err := getTTSKeyFromHTML()
	if err != nil {
		return "", err
	}

	return CalculateToken(text, key1, key2), nil
}

func getTTSKeyFromHTML() (key1, key2 int64, err error) {
	const tokenURL = "https://translate.google.com"
	resp, err := http.Get(tokenURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	matchList := tokenRegex.FindStringSubmatch(string(body))
	if len(matchList) != 3 {
		return 0, 0, fmt.Errorf("HTML parse error")
	}
	part1, _ := strconv.ParseInt(matchList[1], 10, 64)
	part2, _ := strconv.ParseInt(matchList[2], 10, 64)

	return part1, part2, nil
}

// CalculateToken caluculates token from tts text and seed keys.
func CalculateToken(text string, key1, key2 int64) string {
	const salt1 = "+-a^+6"
	const salt2 = "+-3^+b+-f"

	a := key1
	for _, v := range []byte(text) {
		a += int64(v)
		a = workToken(a, salt1)
	}
	a = workToken(a, salt2)

	a ^= key2
	if a < 0 {
		a = (a & 2147483647) + 2147483648
	}
	a %= 1E6
	return fmt.Sprintf("%d.%d", a, a^key1)
}

func workToken(a int64, seed string) int64 {
	for i, max := 0, len(seed)-2; i < max; i += 3 {
		charByte := seed[i+2]
		char := string(charByte)

		var d int64
		switch {
		case char >= "a":
			d = int64(charByte) - 87
		default:
			d, _ = strconv.ParseInt(char, 10, 64)
		}

		switch {
		case string(seed[i+1]) == "+":
			d = rshift(a, d)
		default:
			d = lshift(a, d)
		}

		switch {
		case string(seed[i]) == "+":
			a = (a + d) & 4294967295
		default:
			a ^= d
		}
	}
	return a
}

func rshift(val, n int64) int64 {
	if val >= 0 {
		return val >> uint64(n)
	}
	return (val + 0x100000000) >> uint64(n)
}

func lshift(val, n int64) int64 {
	return val << uint64(n)
}
