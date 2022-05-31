package power

import "github.com/TheAlgorithms/Go/constraints"

// IterativePower is iterative O(logn) function for pow(x, y)
func IterativePower[T constraints.Integer](n, power T) T {
	var res T = 1

	for power > 0 {
		// even => 0, odd => 1, calculate only odd
		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1 // modulus, odd number => odd result, 1 >> 1 yields 0
		n *= n             // multiply with itself
	}
	return res
}

// RecursivePower is recursive O(logn) function for pow(x, y)
func RecursivePower[T constraints.Integer](n, power T) T {
	if power == 0 {
		return 1
	}
	var temp = RecursivePower(n, power/2)
	if power%2 == 0 {
		return temp * temp
	}
	return n * temp * temp
}

// RecursivePower1 is recursive O(n) function for pow(x, y)
func RecursivePower1[T constraints.Integer](n, power T) T {
	if power == 0 {
		return 1
	} else if power%2 == 0 {
		return RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	} else {
		return n * RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	}
}
