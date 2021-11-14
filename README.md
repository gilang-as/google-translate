# google-translate
[![Actions Status](https://github.com/gilang-as/google-translate/actions/workflows/test.yaml/badge.svg)](https://github.com/gilang-as/google-translate/actions)

A **free** and **unlimited** API for Google Translate

Parts of the code are ported from [gtranslate](https://github.com/bregydoc/gtranslate) and [google-translate-api](https://github.com/matheuss/google-translate-api) (also MIT license).

## Features
- Auto language detection
- Spelling correction
- Language correction
- Fast and reliable – it uses the same servers that [translate.google.com](https://translate.google.com) uses

## Install

```
go get github.com/gilang-as/google-translate
```

## API

### Example
```go
package main

import (
	"fmt"
	gtranslate "github.com/gilang-as/google-translate"
)

func main() {
	value := gtranslate.Translate{
		Text: "Halo Dunia",
		//From: "id",
		To: "en",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	}

	dt, err := json.MarshalIndent(translated, "", "\t")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Println(string(dt))
}
```

### Returns an `object`:
- `text` *(string)* – The translated text.
- `pronunciation` *(string)* – The Pronunciation text.
- `from` *(object)*
    - `language` *(object)*
        - `did_you_mean` *(boolean)* - `true` if the API suggest a correction in the source language
        - `iso` *(string)* - The code of the language that the API has recognized in the `text`
    - `text` *(object)*
        - `auto_corrected` *(boolean)* – `true` if the API has auto corrected the `text`
        - `value` *(string)* – The auto corrected `text` or the `text` with suggested corrections
        - `did_you_mean` *(boolean)* – `true` if the API has suggested corrections to the `text`

## License

MIT © [Gilang Adi S](https://github.com/gilang-as)