package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Zoom struct {
	ViewPort rl.Rectangle
	Scale    float32
	Step     float32
}

var ZoomDefault = Zoom{
	ViewPort: rl.Rectangle{X: 0, Y: 0, Width: 200, Height: 100},
	Scale:    1.0,
	Step:     .5,
}

func Draw(tex rl.Texture2D) {}
