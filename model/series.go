package model

type Series struct {
	values map[string]float64
}

func (series *Series) Set(dimension string, value float64) {
	if series.values == nil {
		series.values = make(map[string]float64, 0)
	}

	series.values[dimension] = value
}

// GetDimensions returns the current Series's dimensions
func (series *Series) GetDimensions() []string {
	dimensions := make([]string, len(series.values))
	i := 0
	for d := range series.values {
		dimensions[i] = d
		i++
	}

	return dimensions
}