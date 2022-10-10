package main

func delta(old, new int) int { return new - old } // new作为变量，函数内部不能再使用内置的new函数了。

func main() {
	delta(100, 1)
}
