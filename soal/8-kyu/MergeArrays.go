// Question :

// You are given two sorted arrays that both only contain integers. Your task is to find a way to merge them into a single one, sorted in asc order. Complete the function mergeArrays(arr1, arr2), where arr1 and arr2 are the original sorted arrays.
// You don't need to worry about validation, since arr1 and arr2 must be arrays with 0 or more Integers. If both arr1 and arr2 are empty, then just return an empty array.
// Note: arr1 and arr2 may be sorted in different orders. Also arr1 and arr2 may have same integers. Remove duplicated in the returned result.
// Examples (input -> output)
// * [1, 2, 3, 4, 5], [6, 7, 8, 9, 10] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// * [1, 3, 5, 7, 9], [10, 8, 6, 4, 2] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// * [1, 3, 5, 7, 9, 11, 12], [1, 2, 3, 4, 5, 10, 12] -> [1, 2, 3, 4, 5, 7, 9, 10, 11, 12]

// Answer :
package soal

import (
	"fmt"
	"sort"
)

func MergeArrays(arr1, arr2 []int) (res []int) {

	num := map[int]int{}

	for _, r := range arr1{
		num[r]++
	}
	for _, r := range arr2{
		num[r]++
	}

	fmt.Println(num)


	for _, i := range arr1{
		if num[i] > 0 {
			res = append(res, i)
			num[i] = 0
		}
	}

	for _, i := range arr2{
		if num[i] > 0 {
			res = append(res, i)
			num[i] = 0
		}
	}

	sort.Ints(res)
	
	return
}