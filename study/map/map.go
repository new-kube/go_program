package main

import "fmt"

func main() {
	hello := make(map[string]int)
	hello["hello"] = 1

	delete(hello, "world")

	fmt.Println("delete map not exist key.")

	var nillMap map[string]int

	// nil map可以使用for遍历，不会崩溃。注意这一点。只是不会命中。
	for k, v := range nillMap {
		fmt.Printf("nill map, k=%s v=%d", k, v)
	}
	fmt.Printf("nill map\n")

	keyRange := map[string]int{
		"hello": 1,
		"world": 2,
	}

	// map遍历只关注key，则不需要写v
	for k := range keyRange {
		fmt.Printf("key: %s\n", k)
	}

	var slice []int
	for _, v := range slice {
		fmt.Printf("%d", v)
	}

	fmt.Printf("slice nil range ok\n")
}
