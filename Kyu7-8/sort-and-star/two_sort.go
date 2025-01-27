// https://www.codewars.com/kata/57cfdf34902f6ba3d300001e/train/go

package sortandstar

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}
