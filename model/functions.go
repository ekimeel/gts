package model

type ValueFunction interface {
	Compute(series *TimeSeries) (float64, error)
}
