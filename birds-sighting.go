package main

import "fmt"

/*
  Birdis sighting problem
  - the following array represents a counting of birds sighting based on their types.
  - types are: 1, 2, 3, 4 or 5.
  [ 1, 2, 4, 4, 4, 2, 3, 4 ]

  In this example we have:
  type  | sights
  1     | 1
  2     | 2
  3     | 1
  4     | 4

  Problem:
    - show the most sighted bird (4 in this example)
    - if two types of birds have same sights, so show the one with the minor type
      - [1,1,2,3,4,4,5] = in this example 1 and 4 have 2 sights so, show 1.

  Solution:
    - using a simple array with determined positions for types is the optimal solution for this kind of challenge.
    - the countingSort part is unecessary, it's just a study case showing that is possible to sort the original array without comparisons

*/

// as math does not have loops, this way is a "mathematical" way to do this
// there is a way more "programatic" using for loops
func countingSortMathWay(histo []int32, orig []int32) []int32 {
	// based on a histogram I make a cumulative counting
	// on each index of histogram, this way we can iterate
	// over this histogram and decrease 1 of each iteration
	// and find the correct place to put the value on out array
	for i := 1; i < len(histo); i++ {
		histo[i] += histo[i-1]
	}

	out := make([]int32, len(orig))

	// from last index of original array to start,
	// I put in the
	var birdType int32
	var correctPosition int32
	for i := len(orig) - 1; i >= 0; i-- {
		birdType = orig[i]
		correctPosition = histo[birdType] - 1
		out[correctPosition] = birdType
		// for each iteration we decrease one, this way we
		// can finde the correct place on next iteration, because
		// the current place is was just used by this itreation
		histo[birdType]--
	}
	return out
}

// this way we don't need the original arr
func countingSortProgramaticWay(histo []int32) []int32 {
	var out []int32
	for i := 0; i < len(histo); i++ {
		for j := int32(0); j < histo[i]; j++ {
			out = append(out, int32(i))
		}
	}
	return out
}

func birds(arr []int32) int32 {
	v := make([]int32, 6)

	for i := 0; i < len(arr); i++ {
		v[arr[i]]++
	}
	c := int32(0)
	l := int32(0)

	for i := int32(1); i <= 5; i++ {
		if v[i] > c {
			l = i
			c = v[i]
		}
	}
	copyArrToSortMath := make([]int32, 0, len(v))
	copyArrToSortProg := make([]int32, 0, len(v))
	copyArrToSortMath = append(copyArrToSortMath, v...)
	copyArrToSortProg = append(copyArrToSortProg, v...)
	// extra step just for study purpose
	mSorted := countingSortMathWay(copyArrToSortMath, arr)
	pSorted := countingSortProgramaticWay(copyArrToSortProg)
	fmt.Println("Original Array: ", arr)
	fmt.Println("\nsorted array by counting sort(extra step): ")
	fmt.Println("math way: ", mSorted)
	fmt.Println("programtic way: ", pSorted)

	return l
}
func main() {
	fmt.Print("\033[H\033[2J")
	arr := []int32{1, 2, 4, 4, 4, 2, 3, 4}
	res := birds(arr)
	fmt.Println("\nchallenge result: ", res)
}
