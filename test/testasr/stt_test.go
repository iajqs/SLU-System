package testasr

import (
	"SLU-System/asr"

	"testing"
	"io/ioutil"
	"encoding/base64"
	"time"
)

func TestTrans(t *testing.T){
	a := func(asr *asr.ASR, speech string, response chan string) {
		for {
			resp, err := asr.Trans(speech)
			if err != nil {
				t.Error(err.Error())
			}
			response <- resp
		}
	}

	
	var asr = asr.New()
	speechPath := "./test_1.pcm"
	body, err := ioutil.ReadFile(speechPath)
	if err != nil {
		t.Error(err.Error())
	}
	speech := base64.StdEncoding.EncodeToString(body)

	t.Log(time.Now())

	response := make(chan string, 2)

	go a(asr, speech, response)
	before := time.Now().Unix()
	for {
		result := <- response
		t.Log(time.Now().Unix(), (time.Now().Unix() - before))
		before = time.Now().Unix()
		t.Logf(result)
		
	}

}

