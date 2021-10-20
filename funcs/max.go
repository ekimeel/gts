package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

//returns the Max of a TimeSeries for a provided dimension, return 0 if the TimeSeries is empty
type Max struct {
	Dimension string
}

func (function Max) EvalTimeSeries(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return math.NaN(), fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	if series.Size() == 0 {
		return math.NaN(), nil
	}

	max := math.SmallestNonzeroFloat64
	for i := 0; i < series.Size(); i++ {
		max = math.Max(max, *series.At(i, index))
	}

	return max, nil
}

//EvalSlice returns the maximum value in the provided slice. If the slice is empty a NaN will be returned.
func (function Max) Eval(values []float64) (float64, error) {
	if len(values) == 0 {
		return math.NaN(), nil
	}

	max := math.SmallestNonzeroFloat64
	for i := 0; i < len(values); i++ {
		max = math.Max(max, values[i])
	}
	return max, nil
}


