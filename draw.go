package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	thickness = 3
	radius    = 25.0
)

func draw(state *State) {
	if state.mouseScroll != 0 {
		scale := state.Scale
		state.Scale = rl.Clamp(state.Scale+state.mouseScroll*state.Step, 1, 20)
		if scale != state.Scale {
			ReframeTexture(state)
		}
	}
	rl.DrawTexturePro(state.tex, state.view, state.dest, rl.Vector2Zero(), 0, rl.White)
	drawCrossHair(state, radius)
}

func drawCrossHair(state *State, radius float32) {
	pos := TextureToScreen(state, state.zoomPosition)
	begin := rl.Vector2Add(pos, rl.Vector2{X: -radius})
	end := rl.Vector2Add(pos, rl.Vector2{X: radius})
	rl.DrawLineBezier(begin, end, thickness, rl.Black)
	begin = rl.Vector2Add(pos, rl.Vector2{Y: -radius})
	end = rl.Vector2Add(pos, rl.Vector2{Y: radius})
	rl.DrawLineBezier(begin, end, thickness, rl.Black)
	rl.DrawCircleLinesV(pos, radius, rl.Red)
}
