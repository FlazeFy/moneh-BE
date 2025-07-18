package modules

import (
	"moneh/modules/admin"
	"moneh/schedulers"
	"time"

	"github.com/robfig/cron"
)

func SetUpScheduler(
	adminService admin.AdminService) {

	// Initialize Scheduler
	houseKeepingScheduler := schedulers.NewHouseKeepingScheduler(adminService)

	// Init Scheduler
	c := cron.New()
	Scheduler(c, houseKeepingScheduler)
	c.Start()
	defer c.Stop()
}

func Scheduler(c *cron.Cron, houseKeepingScheduler *schedulers.HouseKeepingScheduler) {
	// For Production
	// Clean Scheduler
	c.AddFunc("0 5 2 * *", houseKeepingScheduler.SchedulerMonthlyLog)

	// For Development
	go func() {
		time.Sleep(5 * time.Second)

		// House Keeping Scheduler
		houseKeepingScheduler.SchedulerMonthlyLog()
	}()
}
