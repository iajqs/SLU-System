package asr

import (
    // "github.com/sirupsen/logrus"
    "github.com/shiguanghuxian/txai" // 引入sdk
)

var (
    appID  = "2159987778"
    appKey = "e51xT4Si6hBfi5hN"
    format = 1
    rate   = 16000
    txAi = txai.New(appID, appKey, false)
)

func (asr *ASR) Trans(speech string) (string, error) {
    response, err := txAi.AaiAsrForBase64(speech, format, rate)
    if err != nil {
        return "", err
    }

    return response.Data.Text, nil
}