/**
 * Created by cks
 * Date: 2020-11-25
 * Time: 15:48
 */

package slu

import (
	"errors"
	"regexp"
)

var (
	key2Answer = map[string]string {
		"你好": "你好",
		"你今年几岁啦": "1岁",
	}
)

func ReluAnswer(question string) (string, error) {
	for key, value := range key2Answer {
		reg1 := regexp.MustCompile(key)
		if reg1 == nil {
			continue
		}
		return value, nil
	}
	return "", errors.New("no answer")
}