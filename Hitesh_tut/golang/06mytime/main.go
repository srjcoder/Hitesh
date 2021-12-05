package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This tutorial file try is about time")

	// present time
	presentTime := time.Now()
	fmt.Println("Present time: ", presentTime)

	// present time with format
	fmt.Println("present time with format: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// date
	createDate := time.Date(2021, time.October, 26, 23, 59, 0, 0, time.Local)
	fmt.Println("print my b'd: ", createDate)
	fmt.Println("b'd with format: ", createDate.Format("01-02-2006 Monday"))
}
