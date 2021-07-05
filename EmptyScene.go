package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type EmptyScene struct {
	scenes       *[20]Scene
	sprite1      Sprite
	currentScene *int
}

func (es *EmptyScene) setQuit(state bool) {
	panic("implement me")
}

func (es *EmptyScene) getQuit() bool {
	panic("implement me")
}

func (es *EmptyScene) init() {
	es.sprite1 = createSprite("texture1.png")
	es.sprite1.create()
}

func (es *EmptyScene) draw(win *pixelgl.Window) {
	es.sprite1.draw(win)
}
func (es *EmptyScene) update(deltaTime float64, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		d := demoscene{}
		changeScene(&d, es.scenes, es.currentScene)
	}
	if win.Pressed(pixelgl.KeyUp) {
		es.sprite1.move(0, 20)
		fmt.Println(es.sprite1.position)
	}
}
