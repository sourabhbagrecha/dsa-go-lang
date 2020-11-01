package array

import (
	"math"
	"sort"
)

// MinimizeTheHeights returns the minimum difference between largest and smallest height after adding/subtracting key from each element once and
func MinimizeTheHeights(arr []int, key int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	sort.Ints(arr)

	ans := arr[n-1] - arr[0]

	small := arr[0] + key
	big := arr[n-1] - key

	if small > big {
		small = small + big
		big = small - big
		small = small - big
	}

	for i := 1; i < n-1; i++ {
		ele := arr[i]

		subtract := ele - key
		add := ele + key

		if subtract >= small || add <= big {
			continue
		}

		if big-subtract <= add-small {
			small = subtract
		} else {
			big = add
		}
	}

	return int(math.Min(float64(ans), float64(big-small)))
}
