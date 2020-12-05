/**
 * Created by cks
 * Date: 2020-11-26
 * Time: 15:49
 */
package slu

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

var (
	urlModel = "http://192.168.1.66:8080/nlp"
	client = &http.Client{}
)

func ModelAnswer(question string) (string, []string, error) {
	requestBody := fmt.Sprintf(`{
		"question": "%s"
	}`, question)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", urlModel, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", []string{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), []string{}, nil
}