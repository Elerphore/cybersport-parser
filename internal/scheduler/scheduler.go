package scheduler

import (
	"elerphore/cybersport-parser/internal/cybersport"
	"fmt"
	"time"

	"atomicgo.dev/schedule"
)

func StatTask() {
	var task = schedule.Every(1*time.Minute, func() bool {
		cybersport.GetNews()

		return true
	})

	fmt.Println(task.IsActive())
	task.Wait()
}
