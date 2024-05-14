package main 

import(
  "fmt"
)

func main(){
  var arr = []int{-2, -3, 4, -1, -2, 1, 5, -3}
  maxSum := MaxSum(arr)
  fmt.Println("max sum: ", maxSum)
}

func MaxSum(arr []int) int {
  // make the first element the maximum one to compare it with other elements
  maxSum := arr[0]
	currSum := 0

	for _, num := range arr {
		currSum += num
    // check if the value of current bigger that the maximum sum that is prisested before
    // if it is bigger then make it the max sum
		if maxSum < currSum {
			maxSum = currSum
		}

    // if current sum a negative number, assign it to zero as it is not needed anymore
    // just search for bigger number or list of numbers in the next iteration(s)
		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}
