package misfunciones

import (
	"strconv"
)

func MakeRange(a, b int) []string {
	r := []string{}

	for i := a; i <= b; i++ {
		r = append(r, strconv.Itoa(i))
	}

	return r
}
