package main

import (
	"fmt"

	"github.com/lkeix/guic"
)

type Handler interface {
	ServeQUIC(r *guic.Request, w guic.ResponseWriter)
}

type handler func(r *guic.Request, w guic.ResponseWriter)

func (h handler) ServeQUIC(r *guic.Request, w guic.ResponseWriter) {
	h(r, w)
}

func main() {
	// リスニングするポートを設定
	addr := ":9999"
	srv, err := guic.NewServer(addr)
	if err != nil {
		panic(err)
	}

	h := handler(func(r *guic.Request, w guic.ResponseWriter) {
		fmt.Println("hogehoge")
	})

	srv.Handler = guic.Handler(h)
	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
