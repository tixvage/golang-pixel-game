package main

import "github.com/faiface/pixel/pixelgl"

type Scene interface {
	init()
	update(dt float64)
	draw(win *pixelgl.Window)
	setQuit(state bool)
	getQuit() bool
}

func changeScene(scene *Scene, scenes *[]*Scene, currentScene int) {
	(*scenes)[currentScene+1] = scene
	(*scenes)[currentScene] = nil

	currentScene++
}
