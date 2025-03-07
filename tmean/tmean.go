package tmean

import (
	"errors"
	"sort"
)

// Creat new function called TMean which can be used later
// create a upperTrim which can accept zerio or more float values
func TMean(data []float64, lowerTrim float64, upperTrim ...float64) (float64, error) {
	// check for empty data slice and give error if it is
	if len(data) == 0 {
		return 0, errors.New("data slice is empty")
	}

	// default to symmetric trimming
	uppertrimval := lowerTrim
	if len(upperTrim) > 0 {
		uppertrimval = upperTrim[0] // Use asymmetric trimming if provided
	}

	// Sort data so we can trip
	sorted_data := make([]float64, len(data))
	// copy data to to avoid modifying the original data slice when sorting
	copy(sorted_data, data)
	sort.Float64s(sorted_data)

	// perform trimming
	l_Count := int(float64(len(sorted_data)) * lowerTrim)
	u_Count := int(float64(len(sorted_data)) * uppertrimval)

	//finalize trim
	trimmed_Data := sorted_data[l_Count : len(sorted_data)-u_Count]

	// Compute mean
	sum := 0.0
	for _, v := range trimmed_Data {
		sum += v
	}
	return sum / float64(len(trimmed_Data)), nil
}
