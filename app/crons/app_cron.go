package crons

import (
	"baseapp/app/database"
	"baseapp/app/models"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type AppCron struct {
}

func NewAppCron() *AppCron {
	return &AppCron{}
}

func (ap *AppCron) RunCron() {
	c := cron.New()

	c.AddFunc("@every 6h", func() {
		fmt.Println("Running cron cleanup expired otp\t--------------")
		database.DB.Unscoped().Where("status = ? AND expired_at < ?", "active", time.Now()).Delete(&models.Otp{})
	})

	c.AddFunc("@every 15m", func() {
		fmt.Println("Running cron update user markfordelte\t--------------")
		database.DB.Model(models.User{}).Where("mark_for_deleted_at < ?", time.Now()).Update("deleted_at", time.Now())
	})

	// Start the cron scheduler
	c.Start()
}
