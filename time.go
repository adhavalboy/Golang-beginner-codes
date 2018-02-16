package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Day())
	fmt.Println(t1.ISOWeek())
	fmt.Println(t1.Zone())
}
