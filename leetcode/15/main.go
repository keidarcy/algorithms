package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	startTime := time.Now()
	// num := []int{}
	// input := getInput()
	// for _, l := range strings.Split(input, "\n") {
	// 	n, _ := strconv.Atoi(l)
	// 	num = append(num, n)
	// }
	results := threeSum(num)
	fmt.Println(results)
	fmt.Println("Time", time.Since(startTime))
}

func threeSum(num []int) [][]int {
	target := 0
	bubbleSort(num)
	results := [][]int{}

	for i := 0; i < len(num)-2; i++ {
		// when it is not equal to the previous value
		if i == 0 || num[i] != num[i-1] {
			start := i + 1
			end := len(num) - 1

			for start < end {
				if num[i]+num[start]+num[end] > target {
					end--
				} else if num[i]+num[start]+num[end] < target {
					start++
				} else {
					results = append(results, []int{num[i], num[start], num[end]})
					start++
					end--

					for start < end && num[start] == num[start-1] {
						start++
					}

					for start < end && num[end] == num[end+1] {
						end--
					}
				}
			}
		}
	}
	return results
}

func bubbleSort(num []int) {
	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num)-i; j++ {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
		}
	}
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}

