package main 

import (
  "testing"
)

type Item struct {
  Arr []int
  Ans int
} 

func TestArrays (t *testing.T) {

  arrays := []Item{
    {Arr: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Ans: 6},
    {Arr: []int{1, 2, 3, 4, 5}, Ans: 15},
    {Arr: []int{-1, -2, -3, -4, -5}, Ans: -1},
    {Arr: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Ans: 6},
    {Arr: []int{-5, -4, -3, -2, -1}, Ans: -1},
    {Arr: []int{5, 4, 3, 2, 1}, Ans: 15},
  }


  for i := 0; i < len(arrays); i++{
    res := MaxSum(arrays[i].Arr)

    if res != arrays[i].Ans {
      t.Fatalf("For array %v, expected: %d, got: %d", arrays[i].Arr, arrays[i].Ans, res)
    }
  }
}
