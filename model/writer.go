package model

type Writer interface {
	Write(series *TimeSeries) error
}
