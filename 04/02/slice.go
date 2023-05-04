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
	extendSummer := summer[:5]
	fmt.Printf("extendSummer=%v, len=%d, cap=%d\n", extendSummer, len(extendSummer), cap(extendSummer))
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

func main() {
	slice_feature()
	empty_and_nil()
}
