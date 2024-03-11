package main

import (
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func main() {
	var restConf rest.RestConf
	conf.MustLoad("hello.yaml", &restConf)
	s, err := rest.NewServer(restConf)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.AddRoute(rest.Route{ // 添加路由
		Method: http.MethodGet,
		Path:   "/hello/world",
		Handler: func(writer http.ResponseWriter, request *http.Request) { // 处理函数
			httpx.OkJson(writer, "Hello World!")
		},
	})

	defer s.Stop()
	s.Start() // 启动服务
}
