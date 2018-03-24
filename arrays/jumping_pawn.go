package arrays

import (
	"fmt"
	"math"
)

/*
Problem Statement
A non-empty zero-indexed array A consisting of N integers is given.
Each element of the array can be treated as a relative pointer to another element in the array ie:
if A[K] = M,  then element A[K] points to element A[K+M].

The array defnes a sequence of jumps of a pawn as follows:

initially, the pawn is located at elementA[0];
on each jump the pawn moves from its current element to the destination element pointed to by the current element; i.e.
if the pawn stands onelement A[K] then it jumps to the element pointed to by A[K];
the pawn may jump forever or may jumpout of the array.

For example, consider the following array A.
A[0] =  2
A[1] =  3
A[2] = -1
A[3] =  1
A[4] =  3

This array defnes the following sequence of jumps of the pawn:
initially, the pawn is located at element A[0];

on the first jump, the pawn moves from A[0] to A[2] because 0 + A[0] = 2;
on the second jump, the pawn movesfrom A[2] to A[1] because 2 + A[2] = 1;
on the third jump, the pawn moves from A[1] to A[4] because 1 + A[1] = 4;

The output should be the number of jumps the pawn takes to exit the array, or -1 if the pawn will remain in the array infinitely.
The output for the above example is 4.
*/

func PawnJumps(array []int) {

	length := len(array)
	var jumps, pos int

	for {

		value := array[pos] + pos
		array[pos] = math.MinInt32
		pos = value
		jumps++

		if pos >= length || pos < 0 {
			fmt.Println(jumps)
			return
		}

		if array[pos] == math.MinInt32 {
			fmt.Printf("-1")
		}
	}

}
