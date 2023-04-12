package main

import (
	"bytes"
	"fmt"
)

func intsToString(ints []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	hello := []int{1, 3, 4, 9, 2, 5, 7, 6, 8}

	fmt.Printf("hello = %s\n", intsToString(hello))
}
