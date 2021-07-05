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
