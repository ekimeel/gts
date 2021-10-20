package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

//returns the Min of a TimeSeries for a provided dimension, return 0 if the TimeSeries is empty
type Min struct {
	Dimension string
}

func (function Min) EvalTimeSeries(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return math.NaN(), fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	if series.Size() == 0 {
		return math.NaN(), nil
	}

	min := math.MaxFloat64
	for i := 0; i < series.Size(); i++ {
		min = math.Min(min, *series.At(i, index))
	}

	return min, nil
}

func (function Min) Eval(values []float64) (float64, error) {
	if len(values) == 0 {
		return math.NaN(), nil
	}

	min := math.MaxFloat64
	for i := 0; i < len(values); i++ {
		min = math.Min(min, values[i])
	}

	return min, nil
}


