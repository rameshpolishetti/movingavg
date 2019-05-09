package cma

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	cma := New()
	avg := cma.Avg()
	if avg != 0 {
		t.Errorf("expected average=%v, got=%v", float64(0), avg)
	}
}

func TestAvg(t *testing.T) {
	cma := New()
	for i := 0; i < 10; i++ {
		cma.AddSample(i)
	}

	avg := cma.Avg() // avg of 0,1...9 = 4.5
	if avg < 4.499 || avg > 4.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}

	for i := 10; i < 100; i++ {
		cma.AddSample(i)
	}
	avg = cma.Avg() // avg of 0,1...99 = 49.5
	if avg < 49.499 || avg > 49.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}
}

func TestAvgInt64(t *testing.T) {
	cma := New()
	for i := 0; i < 10; i++ {
		cma.AddSampleInt64(int64(i))
	}

	avg := cma.Avg() // avg of 0,1...9 = 4.5
	if avg < 4.499 || avg > 4.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}

	for i := 10; i < 100; i++ {
		cma.AddSampleInt64(int64(i))
	}
	avg = cma.Avg() // avg of 0,1...99 = 49.5
	if avg < 49.499 || avg > 49.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}
}

func TestTreadSafeCMA(t *testing.T) {
	ma := New()
	cma := GetThreadSafeCMA(ma)

	// add 10 go routines
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(threadNumber int) {
			for j := 0; j < 10; j++ {
				cma.AddSample(threadNumber)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	avg := cma.Avg() // avg = 4.5
	if avg < 49.499 || avg > 49.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}
}
