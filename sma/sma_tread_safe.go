package sma

import "sync"

// ThreadSafeSMA tread safe version of Simple Moving Average
type ThreadSafeSMA struct {
	mutex sync.RWMutex
	sma   *SMA
}

// TreadSafeSMA returns tread safe version of SMA
func TreadSafeSMA(sma *SMA) *ThreadSafeSMA {
	return &ThreadSafeSMA{
		sma: sma,
	}
}

// AddSample adds a sample to SMA calculator
func (tssma *ThreadSafeSMA) AddSample(s int) {
	tssma.mutex.Lock()
	defer tssma.mutex.Unlock()

	tssma.sma.AddSample(s)
}

// Avg calculates simple moving average and returns
func (tssma *ThreadSafeSMA) Avg() float64 {
	tssma.mutex.Lock()
	defer tssma.mutex.Unlock()

	return tssma.sma.Avg()
}
