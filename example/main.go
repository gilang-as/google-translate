package main

import (
	"fmt"

	gt "gopkg.gilang.dev/google-translate"
	"gopkg.gilang.dev/google-translate/params"
)

func main() {
	value := params.Translate{
		Text: "Hello World!",
		To:   params.INDONESIAN,
	}
	translated, err := gt.TranslateWithParam(value)
	if err != nil {
		panic(err)
	}
	fmt.Println(translated)
}
