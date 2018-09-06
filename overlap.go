package ta

import "errors"

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
	for i, val := range values[timePeriod:] {
		movAvgs = append(movAvgs, runningTotal/timePeriodFloat)
		runningTotal = runningTotal + val - values[i]
	}
	return movAvgs, nil
}
