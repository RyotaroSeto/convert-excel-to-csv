package main

import (
	"time"
)

type ResponseJsonParam struct {
	Status     string `json:"status"`
	Errors     Errors
	AccessTime interface{} `json:"accessTime"`
}

type Errors struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const (
	InputError   string = "パラメータが不正です。"
	DataNotFound string = "紐づくデータが存在しません。"
	NotFound     string = "Not Found."
)

func createErrResponse(status, code string) *ResponseJsonParam {
	result := Errors{Code: code, Msg: errMessage(code)}
	return &ResponseJsonParam{
		Status:     status,
		Errors:     result,
		AccessTime: nowTime(),
	}
}

func errMessage(code string) string {
	var errMessage string
	switch code {
	case "100":
		errMessage = InputError
	case "101":
		errMessage = DataNotFound
	case "102":
		errMessage = NotFound
	}
	return errMessage
}

func nowTime() interface{} {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, loc)
}
