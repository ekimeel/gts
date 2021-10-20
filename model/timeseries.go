package model

import (
	"errors"
	"fmt"
	"time"
)


//A MultivariateTimeSeries
type TimeSeries struct {
	dimensions []string
	times      []int64
	values     [][]float64
}

// NewTimeSeries returns a new TimeSeries with the provided dimensions
func NewTimeSeries(dimensions... string) TimeSeries {
	var result TimeSeries
	result.SetDimensions(dimensions)
	return result
}

// Size returns the number of times in the current TimeSeries
func (ts *TimeSeries) Size() int {
	return len(ts.times)
}

// IsEmpty returns true if the Size() is equal to zero
func (ts *TimeSeries) IsEmpty() bool {
	return ts.Size() == 0
}

// Clear sets values and times equal to nil effectively clearing the current TimeSeries. The Clear func does not
// remove dimensions form the current TimeSeries. Returns true if IsEmpty
func (ts *TimeSeries) Clear() bool {
	ts.values = nil
	ts.times = nil
	return ts.IsEmpty()
}

//Gets a specific value at a time index and dimension
func (ts *TimeSeries) At(timeIndex int, dimension int) *float64 {
	return &ts.values[timeIndex][dimension]
}

func (ts *TimeSeries) CountOfDimensions() int {
	return len(ts.dimensions)
}

// GetTimeAt returns the time at the provided index
func (ts *TimeSeries) GetTimeAt(n int) int64 {
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

// AddSeries adds the provided unix time and Series to the current TimeSeries. Prior to adding the values an
// inspection and mapping is done to ensure dimensions are in order. An error will return if any dimension in
// the Series does not exist current TimeSeries.
//
// This func should not be used when both dimensions of the Series and TimeSeries are known as it introduces additional
// overhead to compute order prior to Add.
func (ts *TimeSeries) AddSeries(time int64, series Series) error {

	if len(ts.dimensions) != len(series.GetDimensions()) {
		return fmt.Errorf("invalid series size, expected %d dimension(s) found: %d", len(ts.dimensions),
			len(series.GetDimensions()))
	}

	values := make([]float64, ts.GetDimensionCount())
	for i, aDim := range series.GetDimensions() {
		j := ts.GetDimensionIndex(aDim)
		if i != j {
			return fmt.Errorf("unknown or out of order dimension: %s", aDim)
		}
		values[j] = series.values[aDim]
	}

	return ts.Add(time, values)


}

func (ts *TimeSeries) Add(time int64, values []float64) error {
	if len(ts.dimensions) != len(values) {
		return fmt.Errorf("invalid dimension size, expected %d found: %d", len(ts.dimensions), len(values))
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

func (ts *TimeSeries) Last() *[]float64 {
	return &ts.values[ts.Size()-1]
}

func (ts *TimeSeries) LastTime() time.Time {
	return time.Unix(ts.times[ts.Size()-1], 0)
}

func (ts *TimeSeries) First() *[]float64 {
	return &ts.values[0]
}

func (ts *TimeSeries) FirstTime() time.Time {
	return time.Unix(ts.times[0], 0)
}


func (ts *TimeSeries) Times() []int64 {
	return ts.times
}

func (ts *TimeSeries) Write(writer Writer) error {
	return writer.Write(ts)
}

func (ts *TimeSeries) Eval(function ValueFunction) (float64, error) {
	return function.EvalTimeSeries(ts)
}


func (ts *TimeSeries) Transform(t Transformation) (TimeSeries, error) {
	return t.Transform(ts)
}

// Map preforms a provided func on each row and adds it to the provided TimeSeries
func (ts *TimeSeries) Map(into *TimeSeries, mapper func(time int64, values []float64) []float64) error {
	for i, row := range ts.values {
		time := ts.GetTimeAt(i)
		series := mapper(time, row)
		err := into.Add(time, series)
		if err != nil {
			return err
		}
	}

	return nil
}

//Filters the current TimeSeries and returns a new one based on the result of the test
func (ts *TimeSeries) Filter(test func(time int64, values []float64) bool) (TimeSeries, error) {
	var filtered TimeSeries
	filtered.SetDimensions(ts.dimensions)

	for i, row := range ts.values {
		time := ts.GetTimeAt(i)
		if test(time, row) == true {
			err := filtered.Add(time, row)
			if err != nil {
				return filtered, err
			}
		}
	}
	return filtered, nil
}


