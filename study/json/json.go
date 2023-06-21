package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	A int     `json:"a,omitempty"`
	B string  `json:"b,omitempty"` // 对应的反解析字符串中无该字段，也能解析成功。
	C float64 `json:"c,omitempty"`
}

var myStruct = `
{
	"a": 1
}
`

func test() {

	data := []byte(myStruct)

	x := MyStruct{}
	if err := json.Unmarshal(data, &x); err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}

	fmt.Printf("succeed, result=%+v\n", x)

	// json: 添加默认值，如果不传该字段的话。
	// http://main.net.cn/faq/program-language/go/how-to-specify-default-values-when-parsing-json-in-go/
	x = MyStruct{
		B: "default b",
		C: 3.1415926,
	}

	if err := json.Unmarshal(data, &x); err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}

	fmt.Printf("succeed, result=%+v\n", x)

}

type World struct {
	C int    `json:"c"`
	D int    `json:"d"`
	E string `json:"e"`
}

type Hello struct {
	A int   `json:"a"`
	B World `json:"b,string"`
}

// 嵌套字符串转换为结构体
var hello = `
{
	"a": 1,
	"b": "{\"c\": 2, \"d\": 3, \"e\": \"hello\"}"
}
`

func test_struct_to_string_wrap() {
	x := Hello{}
	if err := json.Unmarshal([]byte(hello), &x); err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}
	fmt.Printf("succeed, result=%+v\n", x)
}

type Hello2 struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func test_struct_to_string_wrap2() {
	x := Hello2{}
	if err := json.Unmarshal([]byte(hello), &x); err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}
	fmt.Printf("succeed, result=%+v\n", x)

	y := World{}
	if err := json.Unmarshal([]byte(x.B), &y); err != nil {
		fmt.Printf("error wrap=%v\n", err)
		return
	}
	fmt.Printf("succeed, wrap result=%+v\n", x)
}

func main() {
	test_struct_to_string_wrap2()
}
