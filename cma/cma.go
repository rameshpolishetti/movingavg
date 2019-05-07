package cma

// CMA Cumulative Moving Average calculator
type CMA struct {
	avg float64
	n   int
}

// New creats new Cumulative Moving Average calculator
func New() *CMA {
	return &CMA{}
}

// AddSample adds a sample to CMA calculator & returns latest average
func (cma *CMA) AddSample(s int) float64 {
	cma.avg = cma.avg + (float64(s)-cma.avg)/float64(cma.n+1)
	cma.n++
	return cma.avg
}

// Avg returns cumulative moving average
func (cma *CMA) Avg() float64 {
	return cma.avg
}
