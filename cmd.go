package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Cmd int

const (
	CmdNone Cmd = iota
	CmdHome
	CmdEnd
	CmdRight
	CmdLeft
	CmdUp
	CmdDown
)

func (cmd Cmd) String() string {
	switch cmd {
	case CmdHome:
		return "Home"
	case CmdEnd:
		return "End"
	case CmdRight:
		return "Right"
	case CmdLeft:
		return "Left"
	case CmdUp:
		return "Up"
	case CmdDown:
		return "Down"
	}
	return "None"
}

func checkKeys() (cmd Cmd) {
	cmd = CmdNone
	switch {
	case rl.IsKeyDown(rl.KeyHome):
		cmd = CmdHome
	case rl.IsKeyDown(rl.KeyEnd):
		cmd = CmdEnd
	case rl.IsKeyDown(rl.KeyLeft):
		cmd = CmdLeft
	case rl.IsKeyDown(rl.KeyRight):
		cmd = CmdRight
	case rl.IsKeyDown(rl.KeyUp):
		cmd = CmdUp
	case rl.IsKeyDown(rl.KeyDown):
		cmd = CmdDown
	default:
	}
	return
}

func checkPad() (cmd Cmd) {
	cmd = CmdNone
	if rl.IsGamepadAvailable(0) {
		switch {
		case rl.IsGamepadButtonDown(0, rl.GamepadButtonLeftFaceLeft):
			cmd = CmdLeft
		case rl.IsGamepadButtonDown(0, rl.GamepadButtonLeftFaceRight):
			cmd = CmdRight
		case rl.IsGamepadButtonDown(0, rl.GamepadButtonLeftFaceDown):
			cmd = CmdDown
		case rl.IsGamepadButtonDown(0, rl.GamepadButtonLeftFaceUp):
			cmd = CmdUp
		default:
		}
	}
	return
}
