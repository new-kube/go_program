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

func main() {

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
