package main

import (
	"golang.org/x/image/colornames"
)

func run() {
	win := createWindow(colornames.Aqua)
	win.init()
	win.loop()
}
func main() {
	runGame(run)
}
