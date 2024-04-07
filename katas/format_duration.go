package katas

import (
	"fmt"
	"strconv"
	"strings"
)

var minute int64 = 60
var hour int64 = 3600
var day int64 = 24 * hour
var year int64 = 365 * day

var mDuration = map[int64]string{
  year: "year",
  day: "day",
  hour: "hour",
  minute: "minute",
  1: "second",
}

var aDuration = []int64{ year, day, hour, minute, 1 }

func FormatDuration(seconds int64) string {
  var r []string
  for _, k := range aDuration {
    var temp string
    seconds, temp = translate(k, seconds)
    if temp != "" {
      r = append(r, temp)
    }
  }
  var s = r[0]
  for i, v := range r[1:] {
    if i < len(r[1:]) - 1 {
      s = s + ", " + v
    } else {
      s = s + " and " + v
    }
  }
  return s
}

func translate(limit, seconds int64) (int64, string) {
  var t string
  if seconds >= limit {
    l := seconds / limit
    t = fmt.Sprintf("%d %s", l, mDuration[limit])
    if l > 1 {
      t = t + "s"
    }
    seconds = seconds % limit
  }
  return seconds, t
}

func FormatDuration_wasmup(seconds int64) string {
	if seconds == 0 {	return "now" }
	m, seconds := seconds/60, seconds%60
	h, m := m/60, m%60
	d, h := h/24, h%24
	y, d := d/365, d%365
	q := queue(y, "year", nil)
	q = queue(d, "day", q)
	q = queue(h, "hour", q)
	q = queue(m, "minute", q)
	return join(queue(seconds, "second", q))
}
func queue(n int64, s string, q []string) []string {
	if n == 1 {		return append(q, "1 "+s)	}
	if n > 1 {		return append(q, strconv.FormatInt(n, 10)+" "+s+"s")	}
	return q
}
func join(q []string) string {
	if len(q) == 0 {		return "now"	}
	if len(q) == 1 {		return q[0]	}
	return strings.Join(append(q[:len(q)-2], q[len(q)-2]+" and "+q[len(q)-1]), ", ")
}