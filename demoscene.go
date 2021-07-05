package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type demoscene struct {
	scenes *[20]Scene
}

func (es *demoscene) setQuit(state bool) {
	panic("implement me")
}

func (es *demoscene) getQuit() bool {
	panic("implement me")
}

func (es *demoscene) init() {
}

func (es *demoscene) draw(win *pixelgl.Window) {
	fmt.Println("bruh")
}
func (es *demoscene) update(deltaTime float64, win *pixelgl.Window) {
	fmt.Println(deltaTime)
}
