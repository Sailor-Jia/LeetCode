package main

import (
	//"../design"
	"sort"
)

func main() {
	nums := []int{10, 2, 10, 9, 1, 6, 8, 9, 2, 8}
	countKDifference(nums, 5)
	print("jzh")
}

// no.1
func twoSumM1(nums []int, target int) []int {
	//using map value to record the index
	var recordMap = make(map[int][]int)
	for idx, num := range nums {
		if value, ok := recordMap[num]; ok {
			value = append(value, idx)
			recordMap[num] = value
		} else {
			recordMap[num] = []int{idx}
		}
	}

	// find the index
	for idx, num := range nums {
		if value, ok := recordMap[target-num]; ok {
			for _, val := range value {
				if idx != val {
					return []int{idx, val}
				}
			}
		}
	}
	return []int{}
}
func twoSum(nums []int, target int) []int {
	// search while initializing map, can resolve the same value problem
	var recordMap = make(map[int]int)
	for idx, num := range nums {
		if value, ok := recordMap[target-num]; ok {
			return []int{idx, value}
		}
		recordMap[num] = idx
	}
	return []int{}
}

// no.167
func twoSumII(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}

	}
	return []int{}
}

//no.170
//twosum.go
//func twoSumIII() {
//	s := design.NewData()
//	s.Add(3)
//	s.Add(1)
//	s.Add(2)
//	println(s.Find(3))
//	println(s.Find(6))
//}

//no.653
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	//In-order traversal
	valueMap := make(map[int]int, 100)
	return InOrder(root, valueMap, k)

}
func InOrder(root *TreeNode, valueMap map[int]int, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := valueMap[k-root.Val]; ok {
		return true
	}
	valueMap[root.Val] = 1
	return InOrder(root.Left, valueMap, k) || InOrder(root.Right, valueMap, k)
}

//no.1099
func twoSumLessThanK(A []int, k int) int {
	sort.Ints(A)
	i := 0
	j := len(A) - 1
	max := -1
	for i < j {
		if A[i]+A[j] >= k {
			j--
		} else {
			if max < (A[i] + A[j]) {
				max = A[i] + A[j]
			}
			i++
		}
	}
	return max
}

//no.2006
func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				count++
			}
		}
	}
	return count
}

//no.560
func subarraySumForce(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
func subarraySumCumulative(nums []int, k int) int {
	cmlt := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		cmlt[i] = sum
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if cmlt[i] == k {
			count++
		}
		for j := 0; j < i; j++ {
			if cmlt[i]-cmlt[j] == k {
				count++
			}
		}
	}
	return count
}
func subarraySum(nums []int, k int) int {
	cmlMap := make(map[int]int, len(nums))
	cmlMap[0] = 1  // sum==k ,also need add
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += cmlMap[sum-k]
		cmlMap[sum]++  // can not swap with previous sentence.
	}
	return count
}
