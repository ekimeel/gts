package readers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/ekimeel/timeseries/model"
	"io"
	"os"
	"strconv"
	"time"
)

type CsvReader struct {
	Path       string
	TimeLayout string
}

func (csvReader *CsvReader) Read() (model.TimeSeries, error) {
	var ts model.TimeSeries

	file, err := os.Open(csvReader.Path)
	if err != nil {
		return ts, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	header, err := reader.Read()
	if err != nil {
		return ts, errors.New("invalid csv format: unreadable header record")
	}

	if len(header) < 1 || len(header) > 99 {
		return ts, errors.New("invalid csv format: csv must have at least 2 columns and no more than 99")
	}

	if header[0] != "timestamp" {
		return ts, errors.New("invalid csv format: `timestamp` must be first column")
	}

	dimensions := header[1:(len(header))]
	ts.SetDimensions(dimensions)

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			return ts, fmt.Errorf("error during csv reading at row [%d] caused by: %s", i, err)
		}

		time, err := time.Parse(time.RFC3339, record[0])
		if err != nil {
			return ts, fmt.Errorf("invalid timestamp format at row [%d] caused by: not RFC3339 format", i)
		}
		var values = make([]float64, len(record)-1)
		for v := 1; v < len(record); v++ {
			value, err := strconv.ParseFloat(record[v], 64)
			if err != nil {
				return ts, fmt.Errorf("invalid value format at row [%d] and column [%d] caused by: not a number", i, v)
			}
			values[v-1] = value
		}
		err = ts.Add(time.Unix(), values)
		if err != nil {
			return ts, fmt.Errorf("could not add row [%d] to timeseries caused by: %s", i, err)
		}
	}

	return ts, nil

}

