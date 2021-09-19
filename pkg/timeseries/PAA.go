package timeseries

import (
	"math"
	"timeseries/pkg/models"
)

type PAA struct {
	AggPtSize    []int
	originalSize int
	TimeSeries   models.TimeSeries
}

func (paa *PAA) GetOriginalSize() int {
	return paa.originalSize
}

func (paa *PAA) setDimensions(dimensions []string) {
	paa.TimeSeries.SetDimensions(dimensions)
}

func (paa *PAA) Size() int {
	return paa.TimeSeries.Size()
}

func (paa *PAA) Add(time int64, values []float64) error {
	return paa.TimeSeries.Add(time, values)
}

func (paa *PAA) GetAggregatePtSize(ptIndex int) int {
	return paa.AggPtSize[ptIndex]
}

func NewPAA(ts *models.TimeSeries, shrunkSize int) (PAA, error) {
		this := PAA{}

		// Initialize private data.
		this.originalSize = ts.Size()
		this.AggPtSize = make([]int, shrunkSize)

		this.TimeSeries = models.TimeSeries{}
		this.setDimensions(ts.GetDimensions())
		reducedPtSize := float64(ts.Size()) / float64(shrunkSize)

		// Variables that keep track of the range of points being averaged into a single point.
		ptToReadFrom := 0
		ptToReadTo := 0

		for ok := true; ok; ok = ptToReadFrom < ts.Size() {

			ptToReadTo = int(math.Round(reducedPtSize*float64(this.Size()+1)) - 1) // determine end of current range
			ptsToRead := ptToReadTo - ptToReadFrom + 1

			// Keep track of the sum of all the values being averaged to create a single point.
			var timeSum int64
			measurementSums := make([]float64, ts.GetDimensionCount());

			// Sum all of the values over the range ptToReadFrom...ptToReadFrom.
			for pt := ptToReadFrom; pt <= ptToReadTo; pt++ {
				currentPoint := ts.GetMeasurementVector(pt)
				timeSum += ts.GetTimeAtNthPoint(pt)

				for dim := 0; dim < ts.GetDimensionCount(); dim++ {
					measurementSums[dim] += currentPoint[dim]
				}
			}

			// Determine the average value over the range ptToReadFrom...ptToReadFrom.
			timeSum = timeSum / int64(ptsToRead)
			for dim := 0; dim < ts.GetDimensionCount(); dim++ {
				measurementSums[dim] = measurementSums[dim] / float64(ptsToRead) // find the average of each measurement

				// Add the computed average value to the aggregate approximation.
				this.AggPtSize[this.Size()] = ptsToRead
				this.Add(timeSum, measurementSums)

				ptToReadFrom = ptToReadTo + 1 // next window of points to average startw where the last window ended
			}
		}

		return this, nil
}
