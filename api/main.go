// Code generated by hertz generator.

package main

import (
	"simple-douyin/api/biz/client"
	"simple-douyin/api/biz/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	// init rpc client
	client.InitClient()
	// init jwt middleware
	middleware.InitJwt()
}

func main() {
	// for non-framework part
	Init()

	h := server.Default()
	register(h)
	h.Spin()
}
