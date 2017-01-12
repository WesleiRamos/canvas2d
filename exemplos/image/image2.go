package main

import "github.com/WesleiRamos/canvas2d"

var context canvas2d.Context
var lisa *canvas2d.Image
var xy float32

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
		args
					image  = lisa
					pos x  = 0
					pos y  = 0
		(optional)  width  = 600
		(optional)	height = 400

		(optional)	coord x1 = 0
		(optional)	coord y1 = 0
		(optional)	coord x2 = xy
		(optional)	coord y2 = xy
	*/
	context.DrawImage(lisa, 0, 0, 600, 400, 0, 0, xy, xy)

	if xy < 1 {
		xy += 0.005
	}

	if xy > 1 {
		xy = 1
	}
}
