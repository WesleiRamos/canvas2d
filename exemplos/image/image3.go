package main

import "github.com/WesleiRamos/canvas2d"

var context canvas2d.Context
var lisa *canvas2d.Image

func main() {
	canvas := canvas2d.NewCanvas(600, 400, "IMAGEM")
	context = canvas.GetContext()

	canvas.SetSwapInterval(1)
	canvas.SetResizable(false)
	canvas.SetLoadResources(loadResources)
	canvas.SetLoopFunc(Loop)

	canvas.Show()
}

func loadResources() {
	context.Background(canvas2d.NewColor(255, 255, 255))
	lisa = canvas2d.LoadImage("../res/lisa.jpg")
}

func Loop() {
	/*
		image size 234 x 206
					w     h
		args
					image  = lisa
					pos x  = 0
					pos y  = 0

		(optional)	coord x1 = 0.444444 (104 / 234)
		(optional)	coord y1 = 0.048543 (10  / 206)
		(optional)	coord x2 = 1        (234 / 234)
		(optional)	coord y2 = 0.621359 (128 / 206)
	*/
	context.DrawImage(lisa, 183, 97, 0.444444, 0.048543, 1, 0.621359)
}
