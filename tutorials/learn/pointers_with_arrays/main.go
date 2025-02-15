package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf


func multiplyArrVals(arr *[4]int) {
	for i:=0; i<4; i++ {
		arr[i] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	// for i:=0; i<int(numSize); i++ {
	// 	sum += nums[i]
	// }

	for _, val := range nums {
		sum += val
	}

	return sum / numSize
}

func main() {
	pArr := [4]int{1, 2, 3, 4}
	multiplyArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	pf("Average : %.3f\n", getAverage(iSlice...))
}