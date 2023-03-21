package main

import (
	"aliyun/handler"
)

var Ecs2Ip = ""
var Ecs2Port = 80
var Ecs2Path = "/ping"

func main() {
	// ECS1
	r := handler.TestSimpleApi(Ecs2Ip, Ecs2Port, Ecs2Path)

	// ECS2
	//r := handler.TestSimpleApiWithPort("/ping", []byte(`{"ret":"hello world"}`))

	// 该应用在此端口运行
	r.Run(":80")
}
