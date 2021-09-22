package model

type Transformation interface {
	Transform(series *TimeSeries) (TimeSeries, error)
}
