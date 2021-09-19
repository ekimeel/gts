package writers

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"timeseries/pkg/models"
)
const (
	timeStampCol = "timestamp"
	defaultTimeLayout = time.RFC3339
)

type Writer interface {
	Write(series *models.TimeSeries) error
}

func collectHeaders(series *models.TimeSeries) []string  {
	headers := make([]string, series.GetDimensionCount() + 1)
	headers[0] = timeStampCol
	for i, dimension := range series.GetDimensions() {
		headers[i + 1] = dimension
	}
	return headers
}

func collectValuesAt(series *models.TimeSeries, i int, timeLayout string) []string {
	epoch := series.GetTimeAtNthPoint(i)
	measures := series.GetMeasurementVector(i)
	values := make([]string, len(measures) + 1)
	values[0] = time.Unix(epoch, 0).Format(timeLayout)
	for j := 0; j < len(measures); j++ {
		values[j+1] = fmt.Sprintf("%f", measures[j])
	}

	return values
}

/* csv */
type CsvWriter struct {
	Path string
	TimeLayout string
}

func (csvWriter *CsvWriter) Write(series *models.TimeSeries) error {
	file, err := os.Create(csvWriter.Path)
	if err != nil {
		return fmt.Errorf("cannot write to file caused by: %s", err)
	}
	defer file.Close()

	timeLayout := csvWriter.TimeLayout
	if len(timeLayout) == 0 {
		timeLayout = defaultTimeLayout
	}

	/* header */
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := collectHeaders(series)

	err = writer.Write(header)
	if err != nil {
		return fmt.Errorf("failed to write header: %s", err)
	}

	for i := 0; i < series.Size(); i++ {
		values := collectValuesAt(series, i, timeLayout)

		err := writer.Write(values)
		if err != nil {
			return fmt.Errorf("error writing values at row [%d]: %s", i, err)
		}
	}

	return nil
}