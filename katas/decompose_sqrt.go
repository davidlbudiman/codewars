package katas

import "math"

func Decompose(n int64) (res []int64) {     
    pown := n * n
    for i := n-1; i > 1 && len(res) == 0; i-- {
      res = decomposeWithK(pown, i)
    }
    return res
}

func decomposeWithK(pown, xk int64) []int64 {
  var res []int64
  for i:=xk; i > 0; {
    pown = pown - i*i
    temp := res
    res = []int64{i}
    res = append(res, temp...)
    i = int64(math.Floor(math.Sqrt(float64(pown))))
  }
  if pown == 0 && noDuplicates(res) {
    return res
  }
  return []int64{}
}

func noDuplicates(stack []int64) bool {
  m := map[int64]int{}
  for _, v := range stack {
    _, ok := m[v]
    if ok {
      return false
    } else {
      m[v] = 1
    }
  }
  return true
}

func loop(s int64, i int64) []int64 {
  if s < 0 {return nil}
  if s == 0 {return []int64{}}
  var j int64 = i - 1
  for ; j > 0; j-- {
      var sub = loop(s - j * j, j)
      if sub != nil {return append(sub, []int64{j}...)}
  }
  return nil
}
func Decompose_g964(n int64) []int64 {     
  return loop(n*n, n)
}
