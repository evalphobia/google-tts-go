Google TTS (Text-To-Speech) for golang
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/google-tts-go?status.svg
[2]: https://godoc.org/github.com/evalphobia/google-tts-go
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/google-tts-go.svg
[6]: https://github.com/evalphobia/google-tts-go/releases/latest
[7]: https://travis-ci.org/evalphobia/google-tts-go.svg?branch=master
[8]: https://travis-ci.org/evalphobia/google-tts-go
[9]: https://coveralls.io/repos/evalphobia/google-tts-go/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/google-tts-go?branch=master
[11]: https://codecov.io/github/evalphobia/google-tts-go/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/google-tts-go?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/google-tts-go
[14]: https://goreportcard.com/report/github.com/evalphobia/google-tts-go
[15]: https://img.shields.io/github/downloads/evalphobia/google-tts-go/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/google-tts-go/releases
[17]: https://img.shields.io/github/stars/evalphobia/google-tts-go.svg
[18]: https://github.com/evalphobia/google-tts-go/stargazers


google-tts-go is a golang implementation of the token validation of Google Translate.


# Quick Usage

```go
package main

import (
	"fmt"

	"github.com/evalphobia/google-tts-go/googletts"
)

func main() {
	url, err := googletts.GetTTSURL("Hello world.", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(url) // => https://translate.google.com/translate_tts?client=t&ie=UTF-8&q=Hello%2C+world.&textlen=13&tk=368668.249914&tl=en

	tk, err := googletts.GetTTSToken("Hello world.")
	if err != nil {
		panic(err)
	}
	fmt.Println(url) // => 368668.249914
}
```

# Credit

The algorithm of this library is based on [Boudewijn26/gTTS-token](https://github.com/Boudewijn26/gTTS-token) by [Boudewijn26](https://github.com/Boudewijn26/).
