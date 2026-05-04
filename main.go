package main

import (
	"fmt"
	"net/http"

	setting "example.com/m/pkg"
	"example.com/m/routers"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
