package model

type ValueFunction interface {
	Eval(series *TimeSeries) (float64, error)
}
