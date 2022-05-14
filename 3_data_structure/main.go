package main

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	//problem 1 - arraymerge
	fmt.Println(arrayMerge([]string{"king", "jin"}, []string{"eddie", "steve"}))
	fmt.Println(arrayMerge([]string{"king", "jin", "yoshimitsu"}, []string{"eddie", "jin", "yoshimitsu", "gon"}))
	//problem 2 - angka muncul sekali
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))

	//
	pairSum([]int{1, 2, 3, 4, 6}, 6)
	pairSum([]int{4, 3, 6, 7, 8, 1, 9}, 100)
}

func appendUnique(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			fmt.Print("element exist : ")
			fmt.Println(ele)
			return slice
		}
	}
	return append(slice, i)
}

func arrayMerge(array1 []string, array2 []string) (result []string) {

	for i := range array2 {

		array1 = appendUnique(array1, array2[i])
	}

	result = array1

	return result
}

func munculSekali(nums string) (arr []int) {
	var tmp = make([]int, 5, 10)
	var twotimes = make([]int, 5, 10)
	var onetimes = make([]int, 0, 1)

	for i := 0; i < len(nums); i++ {
		v, err := strconv.Atoi(string(nums[i]))

		if slices.Contains(tmp, v) {
			twotimes = append(twotimes, v)
		} else {
			tmp = append(tmp, v)
		}
		if err != nil {
			return arr
		}
	}

	for i := 0; i < len(nums); i++ {
		v, err := strconv.Atoi(string(nums[i]))
		if !slices.Contains(twotimes, v) {
			onetimes = append(onetimes, v)
		}
		if err != nil {
			return arr
		}
	}
	arr = onetimes
	return arr
}

func pairSum(arr []int, sum int) {
	mapping := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if mapping[sum-arr[i]] == 0 {
			mapping[arr[i]] = i
		} else {
			fmt.Printf("Pair for given sum is (%d, %d).\n", arr[mapping[sum-arr[i]]], arr[i])
			return
		}
	}
	fmt.Println("Pair not found in given array.")
}
