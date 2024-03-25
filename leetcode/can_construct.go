package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	var ransomNoteMap = map[rune]int{}
	var magazineMap = map[rune]int{}
	for _, c := range ransomNote {
		ransomNoteMap[c] = ransomNoteMap[c] + 1
	}
	for _, c := range magazine {
		magazineMap[c] = magazineMap[c] + 1
	}
	for k, v := range ransomNoteMap {
		v2, ok := magazineMap[k]
		if ok && v <= v2 {
			delete(ransomNoteMap, k)
		} else {
			return false
		}
	}
	return len(ransomNoteMap) <= 0
}

func CanConstruct_fahmi45(ransomNote string, magazine string) bool {
	res := make([]int, 26)
	res2 := make([]int, 26)

	for _, v := range ransomNote {
		res[v-'a']++
	}
	for _, c := range magazine {
		res2[c-'a']++
	}
	for i := 0; i < len(res); i++ {
		if res[i] > res2[i] {
			return false
		}
	}
	return true
}
