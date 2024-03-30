package katas

import (
  "math"
)

func SumOfSquares(n uint64) uint64 {
  var min_res uint64
  for i:=uint64(math.Floor(math.Sqrt(float64(n)))); i > 1; i-- {
    x := findSumOfSquares(n, i)
    if x < min_res || min_res == 0 {
      min_res = x
    }
  }
  return min_res
}

func findSumOfSquares(n uint64, i uint64) uint64 {
  var res []uint64
  for i > 0 {
    n = n - i*i
    res = append(res, i)
    i = uint64(math.Floor(math.Sqrt(float64(n))))
  }
  return uint64(len(res))  
}
