package googletts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

// get token seed from HTML.
var tokenRegex = regexp.MustCompile(`TKK=eval.*var a\\x3d([0-9+-]+);var b\\x3d([0-9+-]+);return ([0-9+-]+)\+\\x27.\\x27`)

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

func getTTSKeyFromHTML() (key1, key2 int, err error) {
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
	if len(matchList) != 4 {
		return 0, 0, fmt.Errorf("HTML parse error")
	}
	part1, _ := strconv.Atoi(matchList[1])
	part2, _ := strconv.Atoi(matchList[2])
	part3, _ := strconv.Atoi(matchList[3])

	return part3, part1 + part2, nil
}

// CalculateToken caluculates token from tts text and seed keys.
func CalculateToken(text string, key1, key2 int) string {
	const salt1 = "+-a^+6"
	const salt2 = "+-3^+b+-f"

	a := key1
	for _, v := range []byte(text) {
		a += int(v)
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

func workToken(a int, seed string) int {
	for i, max := 0, len(seed)-2; i < max; i += 3 {
		charByte := seed[i+2]
		char := string(charByte)

		var d int
		switch {
		case char >= "a":
			d = int(charByte) - 87
		default:
			d, _ = strconv.Atoi(char)
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

func rshift(val, n int) int {
	if val >= 0 {
		return val >> uint(n)
	}
	return (val + 0x100000000) >> uint(n)
}

func lshift(val, n int) int {
	return val << uint(n)
}
