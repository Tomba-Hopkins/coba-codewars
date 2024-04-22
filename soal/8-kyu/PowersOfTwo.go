// Question :

// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
// Examples
// n = 0  ==> [1]        # [2^0]
// n = 1  ==> [1, 2]     # [2^0, 2^1]
// n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]

// Answer :
package soal

import "math"

func PowersOfTwo(n int) (res []uint64) {
		
	for i := 0; i < n + 1; i++{
		res = append(res, uint64(math.Pow(2, float64(i))))
	}
	
	return res
	
  }
