package main

import "fmt"
import "github.com/WesleiRamos/canvas2d"

var clock canvas2d.Clock
var canvas canvas2d.Canvas
var context canvas2d.Context

var font *canvas2d.Font

func main() {
	canvas = canvas2d.NewCanvas(500, 400, "FPS")
	context = canvas.GetContext()
	clock = canvas2d.NewClock()

	canvas.SetSwapInterval(1)
	canvas.SetLoadResources(load)
	canvas.SetResizable(false)
	canvas.SetLoopFunc(loop)

	canvas.Show()
}

func load() {
	context.Background(canvas2d.NewColor(255, 255, 255))

	font = canvas2d.LoadFont("../res/04b_19.ttf", 40)
	context.Fill.Font = font
	context.Fill.Style = canvas2d.NewColor(0, 0, 0)
}

func loop() {
	clock.Tick()
	context.Fill.Text(fmt.Sprintf("FPS: %d", clock.FPS()), 20, 100)
	context.Fill.Text(fmt.Sprintf("DELTA TIME: %f", clock.DeltaTime()), 20, 175)
}