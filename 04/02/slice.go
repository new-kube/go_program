package main

import "fmt"

// 重点内容，需要用心学习。

func slice_feature() {
	// 1. 定义：是一个拥有相同类型的可变元素序列。
	// 2. slice和array是紧密相连的(重要)。
	//  	1. slice是一个轻量级的数据结构，
	// 		2. 提供了访问数组子序列的功能，
	// 		3. 而且slice的底层引用了一个数组对象。
	// 3. slice的三个属性: 指针、长度、容量(重要)
	//    	1. 指针: 指向第一个slice元素对应的底层数组元素的地址。
	// 		2. 长度: slice的长度，即slice中元素的个数。
	// 		3. 容量: 底层数组从slice的第一个元素开始，到底层数组的最后一个元素的数量。

	// 例子
	// 底层数组
	months := [...]string{
		1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9] // 两者重合在了6月，这表明了两个slice可以指向同一个底层数组。
	fmt.Printf("Q2=%v, len=%d, cap=%d\n", Q2, len(Q2), cap(Q2))
	fmt.Printf("summer=%v, len=%d, cap=%d\n", summer, len(summer), cap(summer))

	// 注意: 如果 slice的引用超过了被引用对象的容量，会导致宕机(panic)。
	// 例如:
	//   1.
	//fmt.Printf("%s\n", summer[20]) // panic
	//   2.
	//fmt.Printf("%v\n", summer[:20]) // panic

	// 如果slice的引用超过了当前的长度(注意不是容量，容量超过了会panic)，但是没有超过容量，那么slice会自动扩展。
	// 注意，这个扩展，只是新的slice扩展了，原来的slice不会变化。
	extendSummer := summer[:5]
	fmt.Printf("extendSummer=%v, len=%d, cap=%d\n", extendSummer, len(extendSummer), cap(extendSummer))
	fmt.Printf("summer=%v, len=%d, cap=%d\n", summer, len(summer), cap(summer))

	// make创建一个无名数组，并且只能给对应复制的slice使用。
	// 3种创建方式:
	//   1. make([]T, len)
	//   2. make([]T, len, cap)
	//   3. make([]T, cap)[:len]
	// 如下的初始化等价于 make([]int, 5, 10)，即长度为5，容量为10。
	// 很有意思，注意和make([]int, 10)的区别。
	x := make([]int, 10)[:5]
	fmt.Printf("x=%v, len=%d, cap=%d\n", x, len(x), cap(x))
	x = make([]int, 10)
	fmt.Printf("x=%v, len=%d, cap=%d\n", x, len(x), cap(x))
}

// 注意，slice和array的区别，
//  1. slice是一个引用类型，所以在函数传递的时候，是传递的引用。
//  2. array是一个值类型，所以在函数传递的时候，是传递的值。
//
// 这个细节需要非常小心，因为这个细节会导致很多问题。
func slice_func_param(x []int) {
	// 引用，可以改变传入的slice的值。
}

// 注意: nil和空slice是不同的。
//  1. slice的零值是nil，nil的slice没有底层数组，也没有任何元素。
//  2. 一个nil的slice的长度和容量都是0。
//  3. 空slice和nil的slice是不同的。空slice是有底层数组的，但是底层数组没有任何元素。
func empty_and_nil() {
	var s []int // len(s) == 0, s == nil
	fmt.Printf("len=%d, cap=%d, s == nil %t\n", len(s), cap(s), s == nil)
	s = nil // len(s) == 0, s == nil
	fmt.Printf("len=%d, cap=%d, s == nil %t\n", len(s), cap(s), s == nil)
	s = []int{} // len(s) == 0, s != nil
	fmt.Printf("len=%d, cap=%d, s == nil %t\n", len(s), cap(s), s == nil)
	// 注意nil可以转换为对应的slice类型的。
	s = []int(nil) // len(s) == 0, s == nil
	fmt.Printf("len=%d, cap=%d, s == nil %t\n", len(s), cap(s), s == nil)

	// 通常很多函数多做了兼容slice为nil的情况。
	// 注意for循环中如果slice为nil，那么不会进入循环。
	// 兼容方法。
	s = nil
	// 无论是nil还是empty，语义都一样，所以归一处理。
	if len(s) == 0 {
		fmt.Printf("s is nil or empty, will return\n")
	}

	// 如果slice为nil，那么for循环不会进入。这个也做了兼容处理
	for _, v := range s {
		fmt.Printf("v=%d\n", v)
	}

	fmt.Printf("not panic\n")
}

func append_int(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice仍然有增长的空间，扩展slice
		z = x[:zlen]
	} else {
		// slice已经没有增长的空间，分配一个新的底层数组。
		// 为了达到分摊线性复杂性，容量扩展一倍。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 内置的copy函数
	}
	z[len(x)] = y
	return z
}

func slice_append() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append_int(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	/*
		// int的slice底层更新如下的结构，这个结果更多的是复合结构，与其说是引用，不如说是个复合结构。
		// 并且改动的时候，指针指向的是同一个地址，所以可以改变。len和cap也是可以变化的。两者都可以变化
		// 包括ptr变化、len/cap不变。ptr不变、len/cap变化。ptr变化、len/cap变化。三种情况。
		// 底层的append内置函数远比上面的append_int复杂的多。
		type IntSlice struct {
			ptr *int
			len int
			cap int
		}
	*/
}

func main() {
	//slice_feature()
	//empty_and_nil()
	slice_append()
}
