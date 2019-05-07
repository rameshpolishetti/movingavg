package main

import (
	"fmt"

	"github.com/rameshpolishetti/movingavg/cma"
	"github.com/rameshpolishetti/movingavg/sma"
)

func main() {

	// simple moving average
	data, _ := sma.New(5)

	data.AddSample(1)
	data.AddSample(2)
	data.AddSample(3)
	data.AddSample(4)
	data.AddSample(5)
	avg := data.Avg()
	fmt.Printf("Simple moving average: %f \n", avg)

	data.AddSample(6)
	avg = data.Avg()
	fmt.Printf("Simple moving average: %f \n", avg)

	// cumulative moving average
	data2 := cma.New()
	data2.AddSample(1)
	data2.AddSample(2)
	fmt.Printf("Cumulative moving average: %f \n", data2.Avg())
}
