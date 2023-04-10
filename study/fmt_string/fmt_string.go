package main

import "fmt"

// fmt print

type Hello struct {
	A string
	B World
	C *World
	D []*World
	E map[int]*World
}

type World struct {
	X int
	Y string
}

func (w World) String() string {
	return fmt.Sprintf("{X: %d, Y: %s}", w.X, w.Y)
}

func main() {

	wds := make([]*World, 0)
	wd1 := &World{
		X: 101,
		Y: "i am a slice for d",
	}
	wds = append(wds, wd1)

	wdm := make(map[int]*World)
	wd2 := &World{
		X: 102,
		Y: "i am a map for e",
	}
	wdm[1] = wd2

	hello := Hello{
		A: "hello",
		B: World{
			X: 1,
			Y: "i am b",
		},
		C: &World{
			X: 2,
			Y: "i am c",
		},
		D: wds,
		E: wdm,
	}

	fmt.Printf("%v\n", hello)
	fmt.Printf("%+v\n", hello)
	fmt.Printf("%v\n", &hello)
	fmt.Printf("%+v\n", &hello)
}
