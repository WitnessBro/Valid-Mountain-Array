package Valid_Mountain_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	//var nums = []int{1, 2, 3, 2, 1}             // Input array 	// Value to remove
	var nums = []int{3, 2, 1}           // Input array 	// Value to remove
	var expectedFlag = true             // The expected answer with correct length.
	var flag = validMountainArray(nums) // Calls your implementation

	assert.Equal(t, flag, expectedFlag)
}
func validMountainArray(arr []int) bool {
	peak := false
	if len(arr) < 3 || arr[0] > arr[1] {
		return false
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if arr[i] < arr[i+1] && peak == false {
			continue
		}
		if arr[i] < arr[i+1] && peak == true {
			return false
		}
		if arr[i] > arr[i+1] {
			peak = true
		}
	}
	return peak
}
