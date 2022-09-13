package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println(currentTime.Format("15:04:05.000000"))
}
