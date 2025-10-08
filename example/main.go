package main

import (
	"fmt"

	gt "gopkg.gilang.dev/google-translate"
	"gopkg.gilang.dev/google-translate/params"
)

func main() {
	value := params.Translate{
		Text: "这是第一句话。 这是第二句话。",
		From: "zh-cn",
		To:   "en",
	}
	translated, err := gt.TranslateWithParam(value)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", translated.Text)
	// Output: "This is the first sentence."
}
