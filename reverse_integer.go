/****************************************************/
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//
// Input: 123
// Output: 321
// Example 2:
//
// Input: -123
// Output: -321
// Example 3:
//
// Input: 120
// Output: 21
// Note:
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
/****************************************************/

/****************************************************/
// Notes:
// 1. int may be int32 or int64 depending on compiler
// 2. corner case is int32 range
/****************************************************/

package leetcode

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}

	d := 0
	for x > 0 {
		d = d*10 + x%10
		if flag > 0 && d > math.MaxInt32 {
			return 0
		}
		if flag < 0 && d > -1*math.MinInt32 {
			return 0
		}

		x = x / 10
	}

	if d < 0 {
		d = 0
	}
	return flag * d
}
