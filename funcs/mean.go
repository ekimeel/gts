package funcs

import (
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"math"
)

//Mean of a TimeSeries, NaN if empty
type Mean struct {
	Dimension string
}

func (function Mean) EvalTimeSeries(series *model.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return math.NaN(), fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	if series.Size() == 0 {
		return math.NaN(), nil
	}

	sum, err := series.Eval(&Sum{Dimension: function.Dimension})
	if err != nil {
		return math.NaN(), err
	}

	return sum / float64(series.Size()), nil
}

func (function Mean) Eval(values []float64) (float64, error) {
	if len(values) == 0 {
		return math.NaN(), nil
	}

	sum, err := Sum{}.Eval(values)

	if err != nil {
		return math.NaN(), err
	}

	return sum / float64(len(values)), nil


	return 0, nil
}
