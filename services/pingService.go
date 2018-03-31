package services

import (
	"template/dto/response"
	"template/engine"
	"template/model"
)

func Ping(str string) response.Response {
	var ping model.Ping
	ping.Log = str
	engine.GetORM().Create(&ping)
	return response.Response{Header: 200, Body: "hello world " + str}
}
