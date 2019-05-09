package sma

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	sma, err := New(10)
	if err != nil {
		t.Errorf(err.Error())
	}

	avg := sma.Avg()
	if avg != 0 {
		t.Errorf("expected average=%v, got=%v", float64(0), avg)
	}

	_, err = New(-1)
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestAvg(t *testing.T) {
	sma, _ := New(5)
	avg := sma.Avg()
	if avg != 0 {
		t.Errorf("expected average=%v, got=%v", float64(0), avg)
	}

	sma.AddSample(1)
	avg = sma.Avg() // avg = 1
	if avg < 0.999 || avg > 1.001 {
		t.Errorf("expected average=%v, got=%v", float64(1), avg)
	}

	sma.AddSample(3)
	avg = sma.Avg() // avg = (1 + 3) / 2 =2
	if avg < 1.999 || avg > 2.001 {
		t.Errorf("expected average=%v, got=%v", float64(2), avg)
	}

	sma.AddSample(4)
	sma.AddSample(2)
	sma.AddSample(5)
	avg = sma.Avg() // avg = (1+3+4+2+5) / 5 = 3
	if avg < 2.999 || avg > 3.001 {
		t.Errorf("expected average=%v, got=%v", float64(3), avg)
	}

	sma.AddSample(6)
	avg = sma.Avg() // avg = (3+4+2+5+6) / 5 = 4
	if avg < 3.999 || avg > 4.001 {
		t.Errorf("expected average=%v, got=%v", float64(4), avg)
	}
}

func TestAvgInt64(t *testing.T) {
	sma, _ := New(5)
	avg := sma.Avg()
	if avg != 0 {
		t.Errorf("expected average=%v, got=%v", float64(0), avg)
	}

	sma.AddSampleInt64(1)
	avg = sma.Avg() // avg = 1
	if avg < 0.999 || avg > 1.001 {
		t.Errorf("expected average=%v, got=%v", float64(1), avg)
	}

	sma.AddSampleInt64(3)
	avg = sma.Avg() // avg = (1 + 3) / 2 =2
	if avg < 1.999 || avg > 2.001 {
		t.Errorf("expected average=%v, got=%v", float64(2), avg)
	}

	sma.AddSampleInt64(4)
	sma.AddSampleInt64(2)
	sma.AddSampleInt64(5)
	avg = sma.Avg() // avg = (1+3+4+2+5) / 5 = 3
	if avg < 2.999 || avg > 3.001 {
		t.Errorf("expected average=%v, got=%v", float64(3), avg)
	}

	sma.AddSampleInt64(6)
	avg = sma.Avg() // avg = (3+4+2+5+6) / 5 = 4
	if avg < 3.999 || avg > 4.001 {
		t.Errorf("expected average=%v, got=%v", float64(4), avg)
	}
}

func TestTreadSafeSMA(t *testing.T) {
	ma, _ := New(100)
	sma := GetThreadSafeSMA(ma)

	// add 10 go routines
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(threadNumber int) {
			for j := 0; j < 10; j++ {
				sma.AddSample(threadNumber)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	avg := sma.Avg() // avg = 4.5
	if avg < 4.499 || avg > 4.501 {
		t.Errorf("expected average=%v, got=%f", float64(4.5), avg)
	}
}
