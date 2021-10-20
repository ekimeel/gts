package model

type ValueFunction interface {
	EvalTimeSeries(series *TimeSeries) (float64, error)
	Eval(values []float64) (float64, error)
}
