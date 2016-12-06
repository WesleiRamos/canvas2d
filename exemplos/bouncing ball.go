package main

import "canvas2d"

var canvas canvas2d.Canvas
var context canvas2d.Context

var posy int32 = 50
var posx int32 = 50
var updateValueX int32 = 2
var updateValueY int32 = 2

func main() {
	canvas = canvas2d.NewCanvas(600, 400, "ball")
	context = canvas.GetContext()

	canvas.SetFullScreen(true)
	canvas.SetLoopFunc(loop)

	canvas.Show()
}

func draw() {
	context.Background(canvas2d.Color{0.95, 0.95, 0.95})

	context.Fill.Style = canvas2d.Color{1.0, 0, 0}
	context.Fill.Arc(posx, posy, 20)
}

func collision() {
	if posx < 0 || posx > int32(canvas.Width) {
		updateValueX *= -1
	}

	if posy < 0 || posy > int32(canvas.Height) {
		updateValueY *= -1
	}
}

func updatePos() {
	posx += updateValueX
	posy += updateValueY
}

func loop() {
	collision()
	updatePos()
	draw()
}