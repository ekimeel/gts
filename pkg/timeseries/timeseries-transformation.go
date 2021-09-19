package timeseries

type TimeSeriesTransformation interface {
	Transform(series *TimeSeries) (TimeSeries, error)
}
