package model

import "time"

const (
	DefaultTimeLayout = time.RFC3339
)

type Reader interface {
	Read() (TimeSeries, error)
}

