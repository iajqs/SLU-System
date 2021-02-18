package task

import (
    // "github.com/sirupsen/logrus"
    "github.com/shiguanghuxian/txai" // 引入sdk
)

// todo: 这里之后需要修改成config的形式
var (
    appID  = "2159987778"
    appKey = "e51xT4Si6hBfi5hN"
    format = 2
    rate   = 16000
    txAi = txai.New(appID, appKey, false)
)

func (task *Task) Trans(speech string) (string, error) {
    response, err := txAi.AaiAsrForBase64(speech, format, rate)
    if err != nil {
        return "", err
    }

    return response.Data.Text, nil
}