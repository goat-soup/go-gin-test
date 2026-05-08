package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	cronJob "example.com/m/cron"
	"example.com/m/models"
	setting "example.com/m/pkg"
	"example.com/m/pkg/logging"
	"example.com/m/routers"
)

func main() {
	// 启动定时任务
	c := cronJob.Setup()
	defer c.Stop()
	defer models.CloseDB()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// 服务连接
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 记录日志
			fmt.Printf("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // syscall.SIGINT, syscall.SIGTERM
	<-quit
	log.Println("Shutdown Server ...")
	ctx, concel := context.WithTimeout(context.Background(), 5*time.Second)
	defer concel()
	if err := s.Shutdown(ctx); err != nil {
		logging.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
