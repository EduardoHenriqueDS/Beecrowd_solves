package main

import (
	"fmt"
)

var N int

func main() {
	fmt.Scan(&N)

	hours := (N / 60) / 60
	minutes := int((((float64(N) / 60) / 60) - float64(hours)) * 60)
	ss := int((((((float64(N) / 60) / 60) - float64(hours)) * 60) - float64(minutes)) * 60)

	fmt.Printf("%v:%v:%v", hours, minutes, ss)
}
