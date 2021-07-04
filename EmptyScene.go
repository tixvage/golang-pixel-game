package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type EmptyScene struct {
	sprite1 Sprite
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
func (es *EmptyScene) update(deltaTime float64) {
	fmt.Println(deltaTime)
}
