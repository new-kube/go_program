package main

import (
	"fmt"
	"regexp"
)

func test_prefix() {
	str := "zP hello"
	match := `^\s*[zZ][pP].*`

	reg := regexp.MustCompile(match)
	if reg.MatchString(str) {
		fmt.Printf("matched str=%s by reg=%s\n", str, match)
	}

	str = "  \tzp abc"
	if reg.MatchString(str) {
		fmt.Printf("matched str=%s by reg=%s\n", str, match)
	}
}

func main() {
	test_prefix()
}
