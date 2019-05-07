package main

import (
	"fmt"

	"github.com/rameshpolishetti/movingavg/sma"
)

func main() {
	data, _ := sma.New(5)

	data.AddSample(1)
	data.AddSample(2)
	data.AddSample(3)
	data.AddSample(4)
	data.AddSample(5)
	avg := data.Avg()
	fmt.Printf("Average: %f \n", avg)

	data.AddSample(6)
	avg = data.Avg()
	fmt.Printf("Average: %f \n", avg)
}
