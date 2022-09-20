package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
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
	f, err := os.Open("test.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	g, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println(err)
	}

	f2, err := os.Open("123.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	g2, err := png.Decode(f2)
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Create("test2.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	p1 := image.NewPaletted(image.Rect(0, 0, 200, 200), palette.Plan9)

	draw.Draw(p1, p1.Bounds(), g, image.ZP, draw.Src) //添加图片

	p2 := image.NewPaletted(image.Rect(0, 0, 200, 200), palette.Plan9)
	draw.Draw(p2, p2.Bounds(), g2, image.ZP, draw.Src) //添加图片

	g1 := &gif.GIF{
		Image:     []*image.Paletted{p1, p2},
		Delay:     []int{30, 30},
		LoopCount: 0,
	}

	gif.EncodeAll(w, g1)
	gif.EncodeAll(f1, g1)
}
