package katas_test

import (
	. "davidbudiman.xyz/codewars/katas"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fmt"
	"math/rand"
	"strings"
	"time"
)

func solution(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	y := seconds / (3600 * 24 * 365)
	d := seconds / (3600 * 24) % 365
	h := seconds / 3600 % 24
	m := (seconds % 3600) / 60
	s := seconds % 3600 % 60

	parts := []string{}

	if years := format(y, "year"); years != "" {
		parts = append(parts, years)
	}
	if days := format(d, "day"); days != "" {
		parts = append(parts, days)
	}
	if hours := format(h, "hour"); hours != "" {
		parts = append(parts, hours)
	}
	if mins := format(m, "minute"); mins != "" {
		parts = append(parts, mins)
	}
	if secs := format(s, "second"); secs != "" {
		parts = append(parts, secs)
	}

	if len(parts) <= 2 {
		return strings.Join(parts, " and ")
	}

	return fmt.Sprintf("%s and %s", strings.Join(parts[:len(parts)-1], ", "), parts[len(parts)-1])
}

func format(n int64, unit string) string {
	if n < 1 {
		return ""
	} else if n < 2 {
		return fmt.Sprintf("%d %s", n, unit)
	}
	return fmt.Sprintf("%d %ss", n, unit)
}

var _ = Describe("Test Example", func() {
	It("Test small static inputs", func() {
		Expect(FormatDuration(1)).To(Equal("1 second"))
		Expect(FormatDuration(62)).To(Equal("1 minute and 2 seconds"))
		Expect(FormatDuration(120)).To(Equal("2 minutes"))
		Expect(FormatDuration(3600)).To(Equal("1 hour"))
		Expect(FormatDuration(3662)).To(Equal("1 hour, 1 minute and 2 seconds"))
	})

	It("Test larger static inputs", func() {
		Expect(FormatDuration(15731080)).To(Equal("182 days, 1 hour, 44 minutes and 40 seconds"))
		Expect(FormatDuration(132030240)).To(Equal("4 years, 68 days, 3 hours and 4 minutes"))
		Expect(FormatDuration(205851834)).To(Equal("6 years, 192 days, 13 hours, 3 minutes and 54 seconds"))
		Expect(FormatDuration(253374061)).To(Equal("8 years, 12 days, 13 hours, 41 minutes and 1 second"))
		Expect(FormatDuration(242062374)).To(Equal("7 years, 246 days, 15 hours, 32 minutes and 54 seconds"))
		Expect(FormatDuration(101956166)).To(Equal("3 years, 85 days, 1 hour, 9 minutes and 26 seconds"))
		Expect(FormatDuration(33243586)).To(Equal("1 year, 19 days, 18 hours, 19 minutes and 46 seconds"))
	})

	It("Test random inputs", func() {
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 1; i <= 100; i++ {
			var seconds = rand.Int63n(9999999)
			Expect(FormatDuration(seconds)).To(Equal(solution(seconds)))
		}
	})
})
