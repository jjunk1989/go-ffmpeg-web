package main

import (
	"bytes"
	_ "math"
	"strconv"
)

// input seconds
// output 00:00:00
func formatTimeString(in int) string {
	var buf bytes.Buffer

	if h := in / 3600; h < 10 {
		buf.WriteString("0" + strconv.Itoa(h))
	} else {
		buf.WriteString(strconv.Itoa(h))
	}
	buf.WriteString(":")
	lt := in % 3600
	if m := lt / 60; m < 10 {
		buf.WriteString("0" + strconv.Itoa(m))
	} else {
		buf.WriteString(strconv.Itoa(m))
	}
	buf.WriteString(":")
	if s := lt % 60; s < 10 {
		buf.WriteString("0" + strconv.Itoa(s))
	} else {
		buf.WriteString(strconv.Itoa(s))
	}
	return buf.String()
}
