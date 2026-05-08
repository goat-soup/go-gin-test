package cron

import (
	"log"

	"example.com/m/models"
	"github.com/robfig/cron/v3"
)

// Setup 初始化定时任务
func Setup() *cron.Cron {
	c := cron.New()

	// 每天凌晨 3 点执行硬删除清理
	_, err := c.AddFunc("0 3 * * *", func() {
		log.Println("[cron] 开始执行硬删除清理...")
		models.CleanAllTag()
		models.CleanAllArticle()
		log.Println("[cron] 硬删除清理完成")
	})
	if err != nil {
		log.Fatalf("[cron] 注册定时任务失败: %v", err)
	}

	c.Start()
	log.Println("[cron] 定时任务已启动")

	return c
}
