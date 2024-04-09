// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

// Example 1:

// Input: s = "egg", t = "add"
// Output: true

// Example 2:

// Input: s = "foo", t = "bar"
// Output: false

// Example 3:

// Input: s = "paper", t = "title"
// Output: true

// Constraints:

// 1 <= s.length <= 5 * 104
// t.length == s.length
// s and t consist of any valid ascii character.
package leetcode

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms, mt := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(s); i++ {
		cs, okS := ms[s[i]]
		ct, okT := mt[t[i]]
		if !okS && !okT {
			ms[s[i]] = t[i]
			mt[t[i]] = s[i]
		} else if okS && okT && cs == t[i] && ct == s[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsIsomorphic_v2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms, mt := map[rune]rune{}, map[rune]rune{}
	sr, tr := []rune(s), []rune(t)
	for i := 0; i < len(sr); i++ {
		cs, okS := ms[sr[i]]
		ct, okT := mt[tr[i]]
		if !okS && !okT {
			ms[sr[i]] = tr[i]
			mt[tr[i]] = sr[i]
		} else if okS && okT && cs == tr[i] && ct == sr[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsIsomorphic_v1(s string, t string) bool {
	sSequence := createSequence(s)
	tSequence := createSequence(t)
	return sSequence == tSequence
}

func createSequence(s string) string {
	var m, seq, idx = map[rune]int{}, "", 1
	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = idx
			idx++
		}
		seq = seq + string(rune(m[c]))
	}
	return seq
}

func IsIsomorphic_Shivansu_7(s string, t string) bool {
	map1 := make([]int, 128) // Stores frequency of s
	map2 := make([]int, 128) // Stores frequency of t

	for i := 0; i < len(s); i++ {
		sCh := s[i]
		tCh := t[i]

		if map1[sCh] == 0 && map2[tCh] == 0 {
			map1[sCh] = int(tCh)
			map2[tCh] = int(sCh)
		} else if map1[sCh] != int(tCh) || map2[tCh] != int(sCh) {
			return false
		}
	}
	return true
}
