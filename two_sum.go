package leetcode

import "sort"

/**************************************/
//	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//	You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//	Example:
//
//	Given nums = [2, 7, 11, 15], target = 9,
//
//	Because nums[0] + nums[1] = 2 + 7 = 9,
//	return [0, 1].
/**************************************/

/**************************************/
// Notes:
// 1. input/output
// 2. slice order
// 3. may have negtive numbers
/**************************************/

func TwoSum(nums []int, target int) []int {
	type Node struct {
		Idx   int
		Value int
	}

	nodes := make([]Node, len(nums))
	for t := 0; t < len(nums); t++ {
		nodes[t] = Node{t, nums[t]}
	}
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Value < nodes[j].Value })
	var index []int
	var i = 0
	var j = len(nodes) - 1
	for j > i {
		var sum = nodes[i].Value + nodes[j].Value
		if sum == target {
			if nodes[i].Idx > nodes[j].Idx {
				index = []int{nodes[j].Idx, nodes[i].Idx}
			} else {
				index = []int{nodes[i].Idx, nodes[j].Idx}
			}
			break
		} else if sum < target {
			i++
		} else {
			j--
		}

	}
	return index
}
