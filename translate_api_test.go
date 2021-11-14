package google_translate

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestTranslate_v1(t *testing.T) {
	data, err := translateV1("Halo Dunia", "id", "ja")
	if err != nil {
		t.Error(err)
		t.Fail()
	}else{
		prettyJSON, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}
		fmt.Println(string(prettyJSON))
	}
}
