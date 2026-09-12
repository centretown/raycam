package main

import (
	"log"
	"time"

	"github.com/centretown/avcamx"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SleepTime = time.Millisecond * 10
)

func setup() {
	avFlags := avcamx.NewAvFlags()
	exists := avFlags.HasFile()
	if exists {
		avFlags.Load()
	}
	avFlags.Parse()
	avFlags.Print()
	rl.SetTraceLogLevel(rl.LogError)
}

func main() {
	setup()

	rl.InitWindow(1500, 900, "Raylib Window")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	config := avcamx.VideoConfig{
		Path:  "http://192.168.10.7:9000/video1",
		Codec: "MJPG",
		// Width:  3840,
		// Height: 2160,
		FPS: 30,
	}
	cam := avcamx.NewRemoteCam(config.Path)
	err := cam.Open(&config)
	if err != nil {
		log.Fatal(err, "Failed to open camera at: ", config)
	}
	defer cam.Close()
	buf, err := cam.Read()
	if err != nil {
		log.Fatal(err)
	}
	state.img = rl.LoadImageFromMemory(".jpeg", buf, int32(len(buf)))
	if !rl.IsImageValid(state.img) {
		log.Fatal("img invalid")
	}
	state.tex = rl.LoadTextureFromImage(state.img)
	state.view = rl.Rectangle{X: 0, Y: 0,
		Width:  float32(state.tex.Width),
		Height: float32(state.tex.Height)}
	log.Print("texture: w,h", state.tex.Width, state.tex.Height)
	rl.SetTargetFPS(30)
	ch := make(chan []byte)
	go poll(cam, ch)

	for !rl.WindowShouldClose() {
		select {
		case buf := <-ch:
			unloadTexture(&state)
			loadTexture(&state, buf)
		default:
			time.Sleep(SleepTime)
		}

		update(&state)
		rl.BeginDrawing()
		draw(&state)
		rl.EndDrawing()
	}
	unloadTexture(&state)
}

func poll(cam *avcamx.RemoteCam, ch chan []byte) {
	for {
		buf, err := cam.Read()
		if err == nil {
			ch <- buf
		}
		time.Sleep(SleepTime)
	}
}
