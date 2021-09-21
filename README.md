# gts

## Concepts

### Dimensions



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