package katas

func SplitStrings(str string) (res []string) {
  if len(str)%2 == 1 {
    str = str + "_"
  }
  for i := 0; i < len(str); i = i+2 {
    res = append(res, str[i:i+2])
  }
  return
}
