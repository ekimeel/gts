package transformations

import "timeseries/pkg/models"

type TimeSeriesTransformation interface {
	Transform(series *models.TimeSeries) (models.TimeSeries, error)
}

type PAA struct{
	AggPtSize    []int
	originalSize int
}

func (t PAA) Transform (series *models.TimeSeries) (models.TimeSeries, error) {

}