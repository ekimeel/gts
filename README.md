# gts

## Concepts

### Value Functions
Value Functions operate on the values of a TimeSeries. 

```go
	max, err := timeSeries.Eval( Max{ Dimension: "v0" } )
	fmt.Printf("The max value of dimension v0 is %f", max)
```


### Dimensions

### Map (higher-order function)
The TimeSeries contains a map function that applies a given function to each Series in the same order. It is often called apply-to-all when considered in functional form.

```go 
reader := CsvReader{Path: "../testdata/3x1000.csv"}
ts, _ := reader.Read()

result := NewTimeSeries("sum")

//define higher-order function
sumFunc := func(time int64, values []float64) []float64{
    var sum float64
    // sum all columns into one value
    for i := 0; i < len(values); i++ {
        sum += values[i]
    }
    return []float64{sum}
}

// executes higher-order function into result
err := ts.Map(&result, sumFunv)
	
```




### Transforming

### Evaluating

### Filtering
Filtering a TimeSeries is 

Example: filter out values greater than 1 in the first dimension
```
filtered := ts.Filter(func(time int64, values []float64) bool {
    return values[0] > 1
})
```