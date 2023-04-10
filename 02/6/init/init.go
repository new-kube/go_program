package init

import "fmt"

// 注意init函数允许多个，按照先后顺序执行。一个文件中。
func init() {
	fmt.Printf("i am a init() ...\n")
}

func init() {
	fmt.Printf("i am a init(), too ...\n")
}
