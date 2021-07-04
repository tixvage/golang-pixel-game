package main

import "github.com/faiface/pixel/pixelgl"

type Entity interface {
	draw(win *pixelgl.Window)
}
