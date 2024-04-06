package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findRestaurant())
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "Tapioca Express", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := 6
	nums2 := []int{1, 2, 2}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(canJump([]int{1, 2, 3}))
}

func canJump(nums []int) bool {
	finishIdx := len(nums)

	if nums[0] == 1 && finishIdx == 1 || nums[len(nums)-1] == len(nums) {
		return true
	}

	for i := 0; i < finishIdx; i++ {
		if nums[i] == (finishIdx - (i + 1)) {
			return true
		}
	}

	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	if nums1[0] != 0 && len(nums1) != 0 {
//		if len(nums1) == (m + n) {
//			for i := len(nums1) - n; i < (m + n); i++ {
//				nums1[i] = nums2[i-n]
//			}
//		}
//		for i := 0; i < len(nums1); i++ {
//			for j := 0; j < len(nums1)-1; j++ {
//				if nums1[j] > nums1[j+1] {
//					nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
//				}
//			}
//		}
//	} else {
//		nums1[0] = nums2[0]
//	}
//	fmt.Println(nums1)
//}

type kv struct {
	k string
	v int
}

func findRestaurant(list1 []string, list2 []string) []string {
	//minIdx := 0
	sumOfIndexes := map[string]int{}

	for i1, v1 := range list1 {
		for i2, v2 := range list2 {
			if v1 == v2 {
				sumOfIndexes[v1] = i1 + i2
			}
		}
	}

	kvStruct := []kv{}

	for k, v := range sumOfIndexes {
		kvStruct = append(kvStruct, kv{
			k: k,
			v: v,
		})
	}

	sort.Slice(kvStruct, func(i, j int) bool {
		return kvStruct[i].v < kvStruct[j].v
	})

	//res := []string{}
	//
	//for i1 := range kvStruct {
	//	for i2 := range kvStruct {
	//		if kvStruct[i1].k == kvStruct {
	//			sumOfIndexes[v1] = i1 + i2
	//		}
	//	}
	//}
	return []string{"aboba"}
}
