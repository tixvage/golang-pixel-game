package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Entity interface {
	draw(win *pixelgl.Window)
}

type EntityObj struct {
	position pixel.Vec
}

func (e *EntityObj) move(x, y float64) {
	e.position.Y += y
	e.position.X += x
}
