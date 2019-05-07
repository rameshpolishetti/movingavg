package sma

import "fmt"

// SMA simple moving average calculator
type SMA struct {
	window int // number of latest values to calculate the average

	samples      []float64
	index        int // index of most recently set sample. For every add sample new index = (index+1)%window
	windowFilled bool

	avg float64
}

// New creats new Simple Moving Average calculator
func New(window int) (*SMA, error) {
	if window <= 0 {
		return nil, fmt.Errorf("incorrect window[%d] value, it should be > 0", window)
	}
	sma := &SMA{
		window:  window,
		samples: make([]float64, window),
	}

	return sma, nil
}

// AddSample adds a sample to SMA calculator
func (sma *SMA) AddSample(s int) {
	sma.samples[sma.index] = float64(s)
	sma.index = (sma.index + 1) % sma.window

	if !sma.windowFilled && sma.index == 0 {
		sma.windowFilled = true
	}
}

// Avg calculates simple moving average and returns
func (sma *SMA) Avg() float64 {
	if !sma.windowFilled && sma.index == 0 {
		return float64(0)
	}

	var sum = float64(0)
	var samplesLength int
	if sma.windowFilled {
		samplesLength = sma.window
	} else {
		samplesLength = sma.index
	}

	for i := 0; i < samplesLength; i++ {
		sum += sma.samples[i]
	}

	sma.avg = sum / float64(samplesLength)
	return sma.avg
}
