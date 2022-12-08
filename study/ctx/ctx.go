package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	key := "hello"

	a := 111

	ctx2 := context.WithValue(ctx, key, &a)

	if x, ok := ctx2.Value(key).(*int); ok {
		fmt.Printf("x=%v\n", x)
	}

	ctx3 := context.WithValue(ctx2, key, nil)

	if x, ok := ctx3.Value(key).(*int); ok {
		fmt.Printf("nil, x=%v\n", x)
	}

}
