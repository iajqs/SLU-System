/**
 * Created by cks
 * Date: 2020-11-25
 * Time: 16:00
 */
package slu

import (
    // "github.com/sirupsen/logrus"
    "github.com/shiguanghuxian/txai" // 引入sdk
)

// todo: 这里之后需要修改成config的形式
var (
    appID  = "2159987778"
    appKey = "e51xT4Si6hBfi5hN"
    format = 1
    rate   = 16000
    txAi = txai.New(appID, appKey, false)
)

func ChatAnswer(session, question string) (string, error) {
    response, err := txAi.NlpTextchatForText(session, question)
    if err != nil {
        return "", err
    }

    return response.Data.Answer, nil
}