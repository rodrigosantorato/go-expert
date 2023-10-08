package main

import (
	"fmt"
	"time"
)

func main() {
	const shortForm = "2006-Jan-02"
	start, err := time.Parse(shortForm, "2023-Oct-01")
	if err != nil {
		panic(err)
	}
	end, err := time.Parse(shortForm, "2024-Mar-30")
	if err != nil {
		panic(err)
	}

	totalDays := end.Sub(start).Hours() / 24
	elapsed := time.Now().Sub(start).Hours() / 24

	fmt.Printf("Progress: %.2f%%", elapsed/totalDays*100)
}
