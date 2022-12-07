package main

import "moonlay/app"

func main() {
	app.Start()

	// taskScheduler := chrono.NewDefaultTaskScheduler()
	// now := time.Now()
	// startTime := now.Add(time.Minute * 1)

	// task, err := taskScheduler.Schedule(func(ctx context.Context) {
	// 	log.Print("One-Shot Task")
	// }, chrono.WithTime(startTime))

	// if err == nil {
	// 	log.Print("Task has been scheduled successfully.")
	// }

}
