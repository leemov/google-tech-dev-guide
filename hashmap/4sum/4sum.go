package main

import (
	"sort"
	"strconv"
	"strings"
)

//still TLE due to O(n^4) ==> need to optimize to O(n^3)
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	sum := map[string]int{}

	//1st loop
	for i, val1 := range nums {
		//2nd loop
		for j, val2 := range nums {
			if i == j {
				continue
			}
			//3rd loop
			for k, val3 := range nums {
				if i == k || j == k {
					continue
				}
				for l, val4 := range nums {
					if k == l || l == j || l == i {
						continue
					}

					keyArr := []int{val1, val2, val3, val4}
					sort.Ints(keyArr)
					key := strconv.Itoa(keyArr[0]) + "." + strconv.Itoa(keyArr[1]) + "." + strconv.Itoa(keyArr[2]) + "." + strconv.Itoa(keyArr[3])
					sum[key] = val1 + val2 + val3 + val4
				}
			}
		}
	}

	result := [][]int{}

	for k, val := range sum {
		if val == target {
			arr4 := strings.Split(k, ".")
			sum4 := make([]int, 0, 4)
			for _, v := range arr4 {
				val, _ := strconv.Atoi(v)
				sum4 = append(sum4, val)
			}
			result = append(result, sum4)
		}
	}

}
