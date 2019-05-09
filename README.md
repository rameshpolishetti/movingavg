# Moving Average

Moving Average implementation for `Go`.

## Simple moving average
`Simple moving average` is the unweighted `mean` of previous `n` data points (`n` also referred as `window` size).

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

    data.AddSample(6) // It will overwrite first value (i.e. 1)
	avg := data.Avg() // returns 4.0
}
```

### Thread safe usage

```
    data, _ := sma.New(5)
    tsdata := sma.GeThreadSafeSMA(data)
    data.AddSample(5)
    avg := data.Avg() // returns 5.0
```

## Cumulative moving average

`Cumulative moving average` is sum of all data points divide by number of data points.

### Usage

```
import 	"github.com/rameshpolishetti/movingavg/cma"

func main() {
	data := cma.New()
	data.AddSample(1)
	data.AddSample(2)
	avg := data.Avg() // returns 1.5
}
```

### Thread safe usage

```
    data := cma.New()
    tsdata := cma.GeThreadSafeCMA(data)
    data.AddSample(5)
    avg := data.Avg() // returns 5.0
```

## Reference

[https://en.wikipedia.org/wiki/Moving_average](https://en.wikipedia.org/wiki/Moving_average)