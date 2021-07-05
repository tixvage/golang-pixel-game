package main

import "github.com/faiface/pixel/pixelgl"

type Scene interface {
	init()
	update(dt float64, win *pixelgl.Window)
	draw(win *pixelgl.Window)
	setQuit(state bool)
	getQuit() bool
}

func changeScene(scene Scene, scenes *[20]Scene, currentScene *int) {
	(*scenes)[*currentScene+1] = scene
	(*scenes)[*currentScene] = nil

	*currentScene++
}
