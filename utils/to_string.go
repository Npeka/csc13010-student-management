package utils

import (
	"fmt"
	"strings"
)

func ArrayToString[T any](arr []T) string {
	sb := strings.Builder{}
	for _, f := range arr {
		sb.WriteString(fmt.Sprintf("%v, ", f))
	}
	sb.WriteString(fmt.Sprintf("%v", arr[len(arr)-1]))
	return sb.String()
}
