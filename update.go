package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	renderSize    rl.Vector2
	mouseDown     bool
	mousePosition rl.Vector2
	zoomPosition  rl.Vector2
	mouseScroll   float32

	view  rl.Rectangle
	dest  rl.Rectangle
	Scale float32
	Step  float32
	Cmd   Cmd

	tex rl.Texture2D
	img *rl.Image
}

var (
	state = State{
		Scale: 1,
		Step:  .25,
	}
)

func update(state *State) {
	state.renderSize = rl.Vector2{X: float32(rl.GetRenderWidth()), Y: float32(rl.GetRenderHeight())}
	state.dest = rl.Rectangle{X: 0, Y: 0, Width: state.renderSize.X, Height: state.renderSize.Y}
	state.mouseScroll = rl.GetMouseWheelMove()
	lastMousePosition := state.mousePosition
	state.mousePosition = rl.GetMousePosition()

	cmd := checkKeys()
	if cmd != CmdNone {
		log.Println(cmd)
	}
	cmd = checkPad()
	if cmd != CmdNone {
		log.Println(cmd)
	}
	state.Cmd = cmd

	if state.mouseDown {
		if rl.IsMouseButtonUp(rl.MouseButtonLeft) {
			state.mouseDown = false
			state.zoomPosition = ScreenToTexture(state, state.mousePosition)
		} else if lastMousePosition != state.mousePosition {
			log.Print("MousePosition last,current", lastMousePosition, state.mousePosition)
			// state.zoomPosition += rl.Vector2Scale(
			// 	rl.Vector2Subtract(state.mousePosition, lastMousePosition),
		}
	} else if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		state.mouseDown = true
	}
}

func viewCoords(state *State) (rl.Vector2, rl.Vector2) {
	return rl.Vector2{X: state.view.X, Y: state.view.Y},
		rl.Vector2{X: state.view.Width, Y: state.view.Height}
}

func TextureToScreen(state *State, pos rl.Vector2) rl.Vector2 {
	viewStart, viewSize := viewCoords(state)
	rv := rl.Vector2Divide(
		rl.Vector2Subtract(pos, viewStart),
		rl.Vector2Divide(viewSize, state.renderSize))
	return rv
}

func ScreenToTexture(state *State, pos rl.Vector2) rl.Vector2 {
	viewStart, viewSize := viewCoords(state)
	rv := rl.Vector2Add(rl.Vector2Multiply(pos,
		rl.Vector2Divide(viewSize, state.renderSize)), viewStart)
	return rv
}

func ReframeTexture(state *State) {
	width := float32(state.tex.Width) / state.Scale
	height := float32(state.tex.Height) / state.Scale
	x := state.zoomPosition.X - width/2
	y := state.zoomPosition.Y - height/2
	x = rl.Clamp(x, 0, float32(state.tex.Width)-width)
	y = rl.Clamp(y, 0, float32(state.tex.Height)-height)
	state.view = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
}

func unloadTexture(state *State) {
	rl.UnloadImage(state.img)
	rl.UnloadTexture(state.tex)
}

func loadTexture(state *State, buf []byte) {
	state.img = rl.LoadImageFromMemory(".jpeg", buf, int32(len(buf)))
	state.tex = rl.LoadTextureFromImage(state.img)
}
