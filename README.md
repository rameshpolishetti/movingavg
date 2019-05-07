# Moving Average

Moving Average implementation for `Go`.

## Simple moving average
Simple moving average (`SMA`) is the unweighted `mean` of previous `n` data points.

### Usage

```
import 	"github.com/rameshpolishetti/movingavg/sma"

func main() {
	data, _ := sma.New(5) // 5 is the window size
	data.AddSample(1)
	data.AddSample(2)
	data.AddSample(3)
	data.AddSample(4)
	data.AddSample(5)
	avg := data.Avg() // returns 3.0

    data.AddSample(6) // It will overwrite firsh value (i.e. 1)
	avg := data.Avg() // returns 4.0
}
```

### Thread safe usage

```
    data, _ := sma.New(5)
    tsdata := sma.TreadSafeSMA(data)
    data.AddSample(5)
    avg := data.Avg() // returns 5.0
```

## Reference

[https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average](https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average)