package google_translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	HOST = "google.com"
)

type req struct {
	FsId string `json:"f.sid"`
	Bl   string `json:"bl"`
	At   string `json:"at"`
}

func extract(key string, value string) string {
	var regex, err = regexp.Compile(`"` + key + `":".*?"`)

	if err != nil {
		fmt.Println(err.Error())
	}
	var res = regex.FindString(value)
	if res == "" {
		return ""
	}
	replace := strings.Replace(res, `"`+key+`":"`, "", -1)
	return replace[:len(replace)-1]
}

func check() (*req, error) {
	var (
		err     error
		client  = &http.Client{}
		baseUrl = "https://translate." + HOST
	)
	request, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error! Initial Check Request.")
	}
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "en-US,en;q=0.9,id;q=0.8")
	request.Header.Set("sec-ch-ua", ".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	request.Header.Set("sec-fetch-dest", "document")
	request.Header.Set("sec-fetch-user", "?1")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error! Bad Network.")
	}
	defer response.Body.Close()
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error! Parsing Data Check.")
	}

	return &req{
		FsId: extract("FdrFJe", string(raw)),
		Bl:   extract("cfb2h", string(raw)),
		At:   extract("SNlM0e", string(raw)),
	}, nil
}

type TranslateFromLanguage struct {
	DidYouMean bool   `json:"did_you_mean"`
	Iso        string `json:"iso"`
}

type TranslateFromText struct {
	AutoCorrected bool    `json:"auto_corrected"`
	Value         *string `json:"value"`
	DidYouMean    bool    `json:"did_you_mean"`
}

type TranslateFrom struct {
	Language TranslateFromLanguage `json:"language"`
	Text     TranslateFromText     `json:"text"`
}

type Translated struct {
	Text          string        `json:"text"`
	Pronunciation *string       `json:"pronunciation"`
	From          TranslateFrom `json:"from"`
}

func translateV1(text string, from string, to string) (*Translated, error) {
	var (
		rpcId   = "MkEWBc"
		err     error
		client  = &http.Client{}
		param   = url.Values{}
		body    = url.Values{}
		baseUrl = "https://translate." + HOST
	)

	u, err := url.Parse(baseUrl + "/_/TranslateWebserverUi/data/batchexecute")
	if err != nil {
		return nil, fmt.Errorf("Base URL not Valid : %s !", baseUrl)
	}

	checkData, err := check()
	if err != nil {
		return nil, err
	}
	query := map[string]string{
		"rpcids":       rpcId,
		"f.sid":        checkData.FsId,
		"bl":           checkData.Bl,
		"hl":           "en-US",
		"soc-app":      "1",
		"soc-platform": "1",
		"soc-device":   "1",
		"_reqid":       strconv.Itoa(int(math.Floor(100000 + (rand.Float64() * 9000)))),
		"rt":           "c",
	}
	for k, v := range query {
		param.Add(k, v)
	}
	u.RawQuery = param.Encode()

	value := [2]interface{}{
		[4]interface{}{
			text,
			from,
			to,
			true,
		},
		[1]interface{}{
			nil,
		},
	}
	values, err1 := json.Marshal(value)
	if err1 != nil {
		return nil, fmt.Errorf("Error! Parsing data 1 to json.")
	}
	data := [1]interface{}{
		[1]interface{}{
			[4]interface{}{
				rpcId,
				string(values),
				nil,
				"generic",
			},
		},
	}
	fReq, err2 := json.Marshal(data)
	if err2 != nil {
		return nil, fmt.Errorf("Error! Parsing data 2 to json.")
	}
	body.Set("f.req", string(fReq))
	var payload = bytes.NewBufferString(body.Encode())
	request, err := http.NewRequest("POST", u.String(), payload)
	if err != nil {
		return nil, fmt.Errorf("Error! Initial Request.")
	}
	request.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	//request.Header.Set("x-goog-batchexecute-bgr", "[key, null,null,345,29,null,null,0,\"2\" ]")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	request.Header.Set("x-same-domain", "1")
	request.Header.Set("sec-ch-ua-mobile", "?1")
	request.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Mobile Safari/537.36")
	request.Header.Set("sec-ch-ua-platform", "\"Android\"")
	request.Header.Set("accept", "*/*")
	request.Header.Set("origin", "https://translate.google.com")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("accept-language", "en-US,en;q=0.9")
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error! Bad Network.")
	}
	defer response.Body.Close()

	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error! GetCookies Body.")
	}
	res := strings.Split(string(raw)[6:], "\n")
	var resp []interface{}
	err = json.Unmarshal([]byte(res[1]), &resp)
	if err != nil {
		return nil, fmt.Errorf("Error! Parsing response 1.")
	}
	if resp[0].([]interface{})[2] == nil {
		return nil, fmt.Errorf("Error! Request on google translate api isn't working, please check your parameter.")
	}
	var resp2 []interface{}
	err = json.Unmarshal([]byte(resp[0].([]interface{})[2].(string)), &resp2)
	if err != nil {
		return nil, fmt.Errorf("Error! Parsing response 2.")
	}

	// Did you mean & autocorrect
	DidYouMean := false
	DidYouMeanLanguage := false
	AutoCorrected := false
	var AutoCorrectedValue *string
	if resp2[0].([]interface{})[0] == nil {
		if resp2[0].([]interface{})[1] != nil && resp2[0].([]interface{})[1].([]interface{})[0] != nil {
			aaaa := resp2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[1].(string)
			r := regexp.MustCompile(`<.*?>`)
			txt := r.ReplaceAllString(aaaa, "")
			AutoCorrectedValue = &txt
			DidYouMean = true
		}
	} else {
		AutoCorrected = true
		DidYouMeanLanguage = true
		txt := resp2[0].([]interface{})[0].(string)
		AutoCorrectedValue = &txt
	}
	textTo := resp2[1].([]interface{})[0].([]interface{})[0].([]interface{})[5].([]interface{})[0].([]interface{})[0].(string)
	pronunciationfrom := resp2[1].([]interface{})[0].([]interface{})[0].([]interface{})[1]
	textIso := resp2[1].([]interface{})[3].(string)
	var pronunciation *string
	if pronunciationfrom != nil {
		a := pronunciationfrom.(string)
		pronunciation = &a
	} else {
		pronunciation = nil
	}
	return &Translated{
		Text:          textTo,
		Pronunciation: pronunciation,
		From: TranslateFrom{
			Language: TranslateFromLanguage{
				DidYouMean: DidYouMeanLanguage,
				Iso:        textIso,
			},
			Text: TranslateFromText{
				AutoCorrected: AutoCorrected,
				Value:         AutoCorrectedValue,
				DidYouMean:    DidYouMean,
			},
		},
	}, nil
}
