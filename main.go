package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	frequency := make(map[int]int)
	var result []int

	for _, value := range a {
		if _, err := frequency[value]; !err {
			frequency[value] = 1
		} else {
			frequency[value] += 1
		}
	}
	for _, value := range b {
		if count, err := frequency[value]; err && count > 0 {
			frequency[value] -= 1
			result = append(result, value)
		}
	}
	return result
}

func main() {
	arr := [][]int{}

	a := []int{5, 3, 1, 2, 7, 12}
	b := []int{6, 2, 12, 4, 5}
	c := []int{1, 1, 1, 2}
	d := []int{1, 1, 1, 1}

	arr = append(arr, a, b, c, d)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("пересечение массива № %v - %v и массива № %v - %v: %v\n", i, arr[i], i+1, arr[i+1],
			intersection(arr[i], arr[i+1]))
	}
}
