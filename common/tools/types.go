package tools

import (
	"strconv"
	"strings"
)

//ToNumberArr "1,2,3,4" to [1,2,3,4]
func ToNumberArr(s string) []int {
	a := strings.Split(s, ",")
	var b []int
	for i := 0; i < len(a); i++ {
		e, _ := strconv.Atoi(a[i])
		b = append(b, e)
	}
	return b
}
