package main

import "github.com/WesleiRamos/canvas2d"

var canvas canvas2d.Canvas
var context canvas2d.Context

var posy float32 = 50
var posx float32 = 50
var updateValueX float32 = 5
var updateValueY float32 = 5

func main() {
	canvas = canvas2d.NewCanvas(600, 400, "ball")
	context = canvas.GetContext()

	canvas.SetSwapInterval(1)
	canvas.SetFullScreen(true)
	canvas.SetLoopFunc(loop)

	canvas.Show()
}

func draw() {
	context.Background(canvas2d.NewColor(245, 245, 245))

	context.Fill.Style = canvas2d.NewColor(0, 0, 255)
	context.Fill.Circle(posx, posy, 20)
}

func collision() {
	if posx < 0 || posx > float32(canvas.Width) {
		updateValueX *= -1
	}

	if posy < 0 || posy > float32(canvas.Height) {
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