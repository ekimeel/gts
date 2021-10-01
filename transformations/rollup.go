package transformations

import (
	"errors"
	"github.com/ekimeel/timeseries/model"
	"math"
)

//Piecewise Aggregate Approximation (RollUp)
type RollUp struct {
	ShrinkSize int
}

func (t RollUp) Transform(series *model.TimeSeries) (model.TimeSeries, error) {
	var paa model.TimeSeries

	if t.ShrinkSize >= series.Size() {
		return paa, errors.New("cannot preform a Piecewise Aggregate Approximation (RollUp) to a size larger than " +
			"the source time series")
	}

	paa.SetDimensions(series.GetDimensions())
	reducedPtSize := float64(series.Size()) / float64(t.ShrinkSize)

	readFrom := 0
	readTo := 0

	for ok := true; ok; ok = readFrom < series.Size() {

		readTo = int(math.Round(reducedPtSize*float64(paa.Size()+1)) - 1)
		ptsToRead := readTo - readFrom + 1

		var timeSum int64
		measurementSums := make([]float64, series.GetDimensionCount())

		for pt := readFrom; pt <= readTo; pt++ {
			currentPoint := series.GetMeasurementVector(pt)
			timeSum += series.GetTimeAtNthPoint(pt)

			for dim := 0; dim < series.GetDimensionCount(); dim++ {
				measurementSums[dim] += currentPoint[dim]
			}
		}

		timeSum = timeSum / int64(ptsToRead)
		for dim := 0; dim < series.GetDimensionCount(); dim++ {
			measurementSums[dim] = measurementSums[dim] / float64(ptsToRead)
			paa.Add(timeSum, measurementSums)

			readFrom = readTo + 1
		}
	}

	return paa, nil
}
