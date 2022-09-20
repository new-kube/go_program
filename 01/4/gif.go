package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/display", display)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func display(w http.ResponseWriter, q *http.Request) {

	f1, err := os.Create("test.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	p1 := image.NewPaletted(image.Rect(0, 0, 110, 110), palette.Plan9)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			p1.Set(50, y, color.RGBA{uint8(x), uint8(y), 255, 255})

		}
	}
	p2 := image.NewPaletted(image.Rect(0, 0, 210, 210), palette.Plan9)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			p2.Set(x, 50, color.RGBA{uint8(x * x % 255), uint8(y * y % 255), 0, 255})

		}
	}
	g1 := &gif.GIF{
		Image:     []*image.Paletted{p1, p2},
		Delay:     []int{30, 30},
		LoopCount: 0,
	}
	gif.EncodeAll(w, g1)  //浏览器显示
	gif.EncodeAll(f1, g1) //保存到文件中
}
