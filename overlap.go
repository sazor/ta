package ta

import (
	"errors"
)

// MAType indicates type of moving average
type MAType uint8

const (
	// SMA - Simple Moving Average
	SMA MAType = iota
	// EMA - Exponential Moving Average
	EMA
	// WMA - Weighted Moving Average
	WMA
	// DEMA - Double Exponential Moving Average
	DEMA
	// TEMA - Triple Exponential Moving Average
	TEMA
	// TRIMA - Triangular Moving Average
	TRIMA
	// KAMA - Kaufman Adaptive Moving Average
	KAMA
	// MAMA - MESA Adaptive Moving Average
	MAMA
	// T3MA - Triple Exponential Moving Average
	T3MA
)

var (
	ErrNegativeTimePeriod = errors.New("time period must be greater than 0")
	ErrTimePeriodTooBig   = errors.New("time period must not be greater than number of values")
)

// SimpleMA
func SimpleMA(values []float64, timePeriod int) ([]float64, error) {
	if timePeriod < 1 {
		return nil, ErrNegativeTimePeriod
	}
	if len(values) < timePeriod {
		return nil, ErrTimePeriodTooBig
	}
	runningTotal := values[0]
	for _, val := range values[1:timePeriod] {
		runningTotal += val
	}
	movAvgs := make([]float64, 0, len(values)-timePeriod+1)
	timePeriodFloat := float64(timePeriod)
	movAvgs = append(movAvgs, runningTotal/timePeriodFloat)
	for i, val := range values[timePeriod:] {
		runningTotal = runningTotal + val - values[i]
		movAvgs = append(movAvgs, runningTotal/timePeriodFloat)
	}
	return movAvgs, nil
}

// ExponentialMA
func ExponentialMA(values []float64, timePeriod int) ([]float64, error) {
	if timePeriod < 1 {
		return nil, ErrNegativeTimePeriod
	}
	if len(values) < timePeriod {
		return nil, ErrTimePeriodTooBig
	}
	k := 2.0 / float64(timePeriod+1)
	runningTotal := values[0]
	for _, val := range values[1:timePeriod] {
		runningTotal += val
	}
	movAvgs := make([]float64, 0, len(values)-timePeriod+1)
	runningTotal /= float64(timePeriod)
	movAvgs = append(movAvgs, runningTotal)
	for _, val := range values[timePeriod:] {
		runningTotal = (val-runningTotal)*k + runningTotal
		movAvgs = append(movAvgs, runningTotal)
	}
	return movAvgs, nil
}

// DoubleExponentialMA
func DoubleExponentialMA(values []float64, timePeriod int) ([]float64, error) {
	if timePeriod < 1 {
		return nil, ErrNegativeTimePeriod
	}
	if len(values)+1 < 2*timePeriod {
		return nil, ErrTimePeriodTooBig
	}
	fEMA, err := ExponentialMA(values, timePeriod)
	if err != nil {
		return nil, err
	}
	sEMA, err := ExponentialMA(fEMA, timePeriod)
	if err != nil {
		return nil, err
	}

	for i := range sEMA {
		sEMA[i] = 2.0*fEMA[i+timePeriod-1] - sEMA[i]
	}

	return sEMA, nil
}
