package tmean

import "testing"

// create test for T_Mean function
func TestT_Mean(t *testing.T) {
	// create slice of values to test
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// run function
	mean, err := TMean(data, 0.1, 0.1)
	// confirm no errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// create varible with expected value
	expected := 5.5

	//confirm we get expcted value
	if mean != expected {
		t.Errorf("Expected %v, got %v", expected, mean)
	}
}
