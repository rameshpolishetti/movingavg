package cma

import "sync"

// ThreadSafeCMA tread safe version of Cumulative Moving Average
type ThreadSafeCMA struct {
	mutex sync.RWMutex
	cma   *CMA
}

// GetThreadSafeCMA returns tread safe version of CMA
func GetThreadSafeCMA(cma *CMA) *ThreadSafeCMA {
	return &ThreadSafeCMA{
		cma: cma,
	}
}

// AddSample adds a sample to CMA calculator & returns latest average
func (tscma *ThreadSafeCMA) AddSample(s int) float64 {
	tscma.mutex.Lock()
	defer tscma.mutex.Unlock()

	return tscma.cma.AddSample(s)
}

// AddSampleInt64 adds a int64 sample to CMA calculator & returns latest average
func (tscma *ThreadSafeCMA) AddSampleInt64(s int64) float64 {
	tscma.mutex.Lock()
	defer tscma.mutex.Unlock()

	return tscma.cma.AddSampleInt64(s)
}

// Avg returns cumulative moving average
func (tscma *ThreadSafeCMA) Avg() float64 {
	tscma.mutex.Lock()
	defer tscma.mutex.Unlock()

	return tscma.cma.Avg()
}
