package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"time"
)

type window struct {
	color        color.RGBA
	_cfg         pixelgl.WindowConfig
	win          *pixelgl.Window
	sprite1      Sprite
	scenes       [20]Scene
	currentScene int
	deltaTime    float64
	last         time.Time
}

func createWindow(rgba color.RGBA) window {
	return window{color: rgba}
}

func (win *window) init() {
	win._cfg = pixelgl.WindowConfig{
		Title:  "go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	var err error
	win.win, err = pixelgl.NewWindow(win._cfg)

	if err != nil {
		panic(err)
	}

	es := EmptyScene{}
	win.scenes[0] = &es
	es.scenes = &win.scenes
	es.init()
	win.currentScene = 0
}

func (win *window) updateDt() {
	win.deltaTime = time.Since(win.last).Seconds()
	win.last = time.Now()
}

func (win *window) render() {
	win.win.Clear(win.color)
	for i, s := range win.scenes {
		if i == win.currentScene {
			if s != nil {
				(s).draw(win.win)
			}
		}
	}
}

func (win *window) update() {
	fmt.Println(int(1 / win.deltaTime))

	for i, s := range win.scenes {
		if i == win.currentScene {
			if s != nil {
				(s).update(win.deltaTime)
			}
		}
	}
}

func (win *window) loop() {
	win.last = time.Now()
	for !win.win.Closed() {

		win.updateDt()
		win.render()
		win.update()

		win.win.Update()
	}
}

func runGame(run func()) {
	pixelgl.Run(run)
}
