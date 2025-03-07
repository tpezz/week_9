package tmean

import (
	"errors"
	"sort"
)

// Creat new function called trimmedMean which can be used later
// create a uppertrid which can accept zerio or more float values
func TMean(data []float64, lowerTrim float64, upperTrim ...float64) (float64, error) {
	// check for empty data slice and give error if it is
	if len(data) == 0 {
		return 0, errors.New("data slice is empty")
	}

	// default to symmetric trimming
	upperTrimVal := lowerTrim
	if len(upperTrim) > 0 {
		upperTrimVal = upperTrim[0] // Use asymmetric trimming if provided
	}

	// Sort data so we can trip
	sortedData := make([]float64, len(data))
	// copy data to to avoid modifying the original data slice when sorting
	copy(sortedData, data)
	sort.Float64s(sortedData)

	// perform trimming
	lowerCount := int(float64(len(sortedData)) * lowerTrim)
	upperCount := int(float64(len(sortedData)) * upperTrimVal)

	//finalize trim
	trimmedData := sortedData[lowerCount : len(sortedData)-upperCount]

	// Compute mean
	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}
	return sum / float64(len(trimmedData)), nil
}
