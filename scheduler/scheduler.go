package scheduler

import (
	"context"
	"moonlay/helper"
	"time"

	"github.com/procyon-projects/chrono"
)

func StartScheduler(callback func(rows [][]string), rows [][]string) {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		callback(rows)
	}, 1*time.Minute)

	helper.PanicIFError(err)

}
