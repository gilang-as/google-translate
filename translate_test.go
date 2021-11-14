package google_translate

import (
	"testing"
)

func TestTranslator(t *testing.T) {
	value := Translate{
		Text: "Halo Dunia",
		//From: "id",
		To: "en",
	}
	translated, err := Translator(value)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}