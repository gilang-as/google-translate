package google_translate

import (
	"github.com/gilang-as/google-translate/params"
	"testing"
)

func TestTranslateWithParam(t *testing.T) {
	value := params.Translate{
		Text: "Halo Dunia",
		//From: "id",
		To: params.ENGLISH,
	}
	translated, err := TranslateWithParam(value)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}

func TestTranslate(t *testing.T) {
	translated, err := Translate("Hello World", params.INDONESIAN)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}

func TestManualTranslate(t *testing.T) {
	translated, err := ManualTranslate("Halo Semuanya", params.INDONESIAN, params.JAVANESE)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}
