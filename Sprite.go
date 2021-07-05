package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/png"
	"os"
)

type Sprite struct {
	EntityObj
	picture pixel.Picture
	sprite  *pixel.Sprite
}

func (spr *Sprite) create() {
	spr.sprite = pixel.NewSprite(spr.picture, spr.picture.Bounds())
}

func (spr *Sprite) draw(win *pixelgl.Window) {
	spr.sprite.Draw(win, pixel.IM)
}

func createSprite(path string) Sprite {
	pix, err := _loadPicture(path)
	if err != nil {
		panic(err)
	}
	spr := Sprite{EntityObj{}, pix, nil}
	return spr
}

func _loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
