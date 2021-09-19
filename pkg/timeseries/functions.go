package timeseries

import (
	"fmt"
	"timeseries/pkg/models"
)

type ValueFunction interface {
	Compute(series *models.TimeSeries) (float64,error)
}

type SumDimension struct {
	Dimension string
}

func (function SumDimension) Compute(series *models.TimeSeries) (float64, error) {
	index := series.GetDimensionIndex(function.Dimension)
	if index < 0 {
		return -1, fmt.Errorf("dimension not found [%s]", function.Dimension)
	}

	sum := 0.0
	for i := 0; i < series.Size(); i++ {
		sum += series.values[i][index]
	}

	return sum, nil
}