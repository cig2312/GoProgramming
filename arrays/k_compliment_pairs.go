package arrays

import (
	"fmt"
)

/*
Problem statement :

Given an array find all such unique pairs of numbers whose sum is equal to 'K'.

eg for an array : [7,2,5,6,5,3,4,0,2] and k = 7

Output would be : (7,0),(0,7),(2,5),(5,2),(3,4),(4,3)
*/

func ComplimentaryPairs(array []int, k int) {

	numbers := make(map[int]int)

	for i := 0; i < len(array); i++ {
		num := array[i]
		target := k - num

		if _, ok := numbers[target]; ok {
			fmt.Printf("(%v, %v)", num, target)
			if num != target {
				fmt.Printf("(%v, %v)", target, num)
			}
		} else {
			numbers[num] = num
		}
	}

}
