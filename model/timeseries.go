package model

import (
	"errors"
	"fmt"
	"github.com/ekimeel/timeseries/transformations"
	"time"
)

//A MultivariateTimeSeries
type TimeSeries struct {
	dimensions []string
	times      []int64
	values     [][]float64
}

func (ts *TimeSeries) Size() int {
	return len(ts.times)
}

//Returns the value at a time and dimension
func (ts *TimeSeries) At(time int, i int) *float64 {
	return &ts.values[time][i]
}

func (ts *TimeSeries) CountOfDimensions() int {
	return len(ts.dimensions)
}

func (ts *TimeSeries) GetTimeAtNthPoint(n int) int64 {
	return ts.times[n]
}

func (ts *TimeSeries) GetDimensions() []string {
	return ts.dimensions
}

func (ts *TimeSeries) SetDimensions(dimensions []string) {
	ts.dimensions = nil
	ts.dimensions = dimensions
}

//Appends the provided dimension to the current TimeSeries
func (ts *TimeSeries) AppendDimension(dimension string, values []float64) (int, error) {
	if ts.HasDimension(dimension) {
		return -1, fmt.Errorf("duplicate dimension [%s]", dimension)
	}

	if ts.Size() != len(values) {
		return -1, errors.New("cannot add dimension of unequal size")
	}

	ts.dimensions = append(ts.dimensions, dimension)
	for i := 0; i < ts.Size(); i++ {
		ts.values[i] = append(ts.values[i], values[i])
	}

	return ts.GetDimensionIndex(dimension), nil
}

// returns true if the current TimeSeries contains the provided dimension
func (ts *TimeSeries) HasDimension(dimension string) bool {
	return ts.GetDimensionIndex(dimension) >= 0
}

func (ts *TimeSeries) GetDimensionCount() int {
	return len(ts.dimensions)
}

func (ts *TimeSeries) GetMeasurementVector(i int) []float64 {
	return ts.values[i]
}

func (ts *TimeSeries) GetDimensionIndex(dimension string) int {
	for i, n := range ts.dimensions {
		if n == dimension {
			return i
		}
	}
	return -1
}

func (ts *TimeSeries) GetDimensionAt(index int) []float64 {
	if index < 0 {
		return nil
	}

	dimension := make([]float64, ts.Size())
	for i := 0; i < ts.Size(); i++ {
		dimension[i] = ts.values[i][index]
	}

	return dimension
}

func (ts *TimeSeries) GetDimension(label string) []float64 {
	index := ts.GetDimensionIndex(label)
	return ts.GetDimensionAt(index)
}

func (ts *TimeSeries) Add(time int64, values []float64) error {
	if len(ts.dimensions) != len(values) {
		return errors.New(fmt.Sprintf("invalid vector size, expected %d found: %d", len(ts.dimensions), len(values)))
	}

	if ts.Size() > 0 && time <= ts.times[len(ts.times)-1] {
		return errors.New("out of order time sequence")
	}

	ts.times = append(ts.times, time)
	ts.values = append(ts.values, values)

	return nil
}

// Convenience for Add(int64, []float)
func (ts *TimeSeries) AddTime(time time.Time, values []float64) error {
	return ts.Add(time.Unix(), values)
}

// Returns the latest known time
func (ts *TimeSeries) LatestTime() time.Time {
	return time.Unix(ts.times[ts.Size()-1], 0)
}

func (ts *TimeSeries) Times() []int64 {
	return ts.times
}

//func (ts *TimeSeries) AddN(time int64, values []float64) error {
func (ts *TimeSeries) Write(writer Writer) error {
	return writer.Write(ts)
}

func (ts *TimeSeries) ComputeValue(function ValueFunction) (float64, error) {
	return function.Compute(ts)
}

func (ts *TimeSeries) Transform(t transformations.Transformation) (TimeSeries, error) {
	return t.Transform(ts)
}

//Filters the current TimeSeries and returns a new one based on the result of the test
func (ts *TimeSeries) Filter(test func(time int64, values []float64) bool) TimeSeries {
	var filtered TimeSeries
	filtered.SetDimensions(ts.dimensions)
	for i, row := range ts.values {
		time := ts.GetTimeAtNthPoint(i)
		if test(time, row) == true {
			filtered.Add(time, row)
		}
	}
	return filtered
}