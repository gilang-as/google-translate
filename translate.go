package google_translate

import (
	"fmt"
	"golang.org/x/text/language"
)

type Translate struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func Translator(value Translate) (*Translated, error) {
	var (
		text string
		from = "auto"
		to   string
	)
	if value.Text == "" {
		return nil, fmt.Errorf("Text Value is required!")
	} else {
		text = value.Text
	}
	if value.To == "" {
		return nil, fmt.Errorf("To Value is required!")
	} else {
		if _, err := language.Parse(value.To); err != nil {
			return nil, fmt.Errorf("To Value is't valid!")
		}
		to = value.To
	}
	if value.From != "" {
		if _, err := language.Parse(value.From); err != nil {
			return nil, fmt.Errorf("From Value is't valid!")
		}
		from = value.From
	}
	return translateV1(text, from, to)
}
