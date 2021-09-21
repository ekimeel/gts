package model

type Reader interface {
	Read() (TimeSeries, error)
}

