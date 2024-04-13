package katas_test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "davidbudiman.xyz/codewars/katas"
)

type testCase struct {
	in  []int
	out string
}

var basicTests = []testCase{
	{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, "-6,-3-1,3-5,7-11,14,15,17-20"},
	{[]int{40, 44, 48, 51, 52, 54, 55, 58, 67, 73}, "40,44,48,51,52,54,55,58,67,73"},
	{[]int{59, 60, 61, 62, 68, 69, 70, 71, 80, 81, 82, 83, 84, 85, 88, 97}, "59-62,68-71,80-85,88,97"},
	{[]int{21, 26, 27, 28, 29, 30, 31, 36, 45, 46, 47}, "21,26-31,36,45-47"},
	{[]int{13, 14, 15, 16, 17, 25, 26, 32, 33, 36, 37, 38, 39, 40}, "13-17,25,26,32,33,36-40"},
}

var _ = Describe("Range Extraction", func() {
	Describe("Basic Tests", func() {
		for i, p := range basicTests {
			seq, out := p.in, p.out
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := RangeExtraction(seq)
				if ans != out {
					fmt.Printf("<LOG::Input>%#v\n", seq)
				}
				Expect(ans).To(match(out))
			})
		}
	})

	Describe("Random Tests", func() {
		for i := 0; i < 50; i += 1 {
			seq := randSeq(randInt(10, 20))
			ref := refRangeExtraction(seq)
			It(fmt.Sprintf("Test #%02d", i+1), func() {
				ans := RangeExtraction(seq)
				if ans != ref {
					fmt.Printf("<LOG::Input>%#v\n", seq)
				}
				Expect(ans).To(match(ref))
			})
		}
	})
})

func match(expected interface{}) types.GomegaMatcher {
	return &rangesMatcher{
		expected: expected,
	}
}

type rangesMatcher struct {
	expected interface{}
}

func (m *rangesMatcher) Match(actual interface{}) (success bool, err error) {
	return actual == m.expected, nil
}
func (m *rangesMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto equal\n\t%#v\n", actual, m.expected)
}
func (m *rangesMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to equal\n\t%#v\n", actual, m.expected)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randInt(a, b int) int { // a <= b
	return r.Intn(b-a) + a
}

func randSeq(size int) []int {
	xs := make([]int, size)
	xs[0] = randInt(-50, 100)
	for i := 1; i < size; i += 1 {
		if r.Float64() < 0.5 {
			xs[i] = xs[i-1] + randInt(1, 10)
		} else {
			xs[i] = xs[i-1] + 1
		}
	}
	return xs
}

func refRangeExtraction(list []int) string {
	n := len(list)
	if n == 0 {
		return ""
	}
	var ss []string
	i := 0
	for j := 1; j < n; j += 1 {
		if list[j-1]+1 != list[j] {
			ss = append(ss, str(list[i:j]))
			i = j
		}
	}
	ss = append(ss, str(list[i:n]))
	return strings.Join(ss, ",")
}

func str(r []int) string {
	n := len(r)
	switch n {
	case 1:
		return fmt.Sprintf("%d", r[0])
	case 2:
		return fmt.Sprintf("%d,%d", r[0], r[1])
	default:
		return fmt.Sprintf("%d-%d", r[0], r[n-1])
	}
}
