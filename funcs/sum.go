package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

type Sum struct {
	Dimension string
}

func (function Sum) EvalTimeSeries(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return -1, fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	sum := 0.0
	for i := 0; i < series.Size(); i++ {
		sum += *series.At(i, index)
	}

	return sum, nil
}

func (function Sum) Eval(values []float64) (float64, error) {
	if len(values) == 0 {
		return math.NaN(), nil
	}

	sum := 0.0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum, nil
}