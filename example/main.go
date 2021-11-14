package main

import (
	"encoding/json"
	"fmt"
	gtranslate "github.com/gilang-as/google-translate"
)

func main()  {
	value := gtranslate.Translate{
		Text: "Halo Dunia",
		//From: "id",
		To: "en",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	}else{
		prettyJSON, err := json.MarshalIndent(translated, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(prettyJSON))
	}
}
