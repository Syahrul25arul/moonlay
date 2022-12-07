package scheduler

import (
	"context"
	"fmt"
	"math"
	"moonlay/helper"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func StartScheduler(totalData int, ctx context.Context, f *excelize.File, domain string, sheetName string, callback func(i int, sheetName string, domain string)) {
	interval, err := strconv.ParseInt(fmt.Sprintf("%.0f", math.Ceil(float64(totalData)/float64(100))), 10, 64)
	helper.PanicIFError(err)

	for i := 0; i < int(interval); i++ {
		time.Sleep(1 * time.Second)
		callback(i, sheetName, domain)
	}

}
