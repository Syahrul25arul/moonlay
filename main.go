package main

import (
	"moonlay/helper"
	"moonlay/scheduler"
)

func main() {
	// sheetnames := []string{"customers", "products", "transactions"}

	// for _, sheetname := range sheetnames {
	// 	rows, err := helper.ReadData(sheetname)
	// 	helper.PanicIFError(err)

	// 	callback := helper.CreateData

	// 	scheduler.StartScheduler(callback, rows)
	// }

	rows, err := helper.ReadData("customers")
	helper.PanicIFError(err)

	callback := helper.CreateData

	scheduler.StartScheduler(callback, rows)

	forever := make(chan string)

	// taskScheduler := chrono.NewDefaultTaskScheduler()

	// _, err = taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
	// 	log.Print("Fixed Rate of 5 seconds")
	// }, 20*time.Second)

	// if err == nil {
	// 	log.Print("Task has been scheduled successfully.")
	// }

	// db := app.NewDB()
	// repository_customer := repository_customer.NewRepositoryCustomer()
	<-forever
}
