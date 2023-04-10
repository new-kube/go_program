package main

import "fmt"

// 创建两个新类型。
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) } // 类型方法集
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func main() {
	c := Celsius(0)

	fmt.Println(c.String())
	fmt.Printf("%s\n", c) // 会调用String()方法。
	fmt.Printf("%v\n", c) // 会调用String()方法
	fmt.Println(c)        // %v 方式，转换 对象
	fmt.Printf("%f\n", c)
	fmt.Println(float64(c))
}
