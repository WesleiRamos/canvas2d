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
		args
			image = lisa
			pos x = 183
			pos y = 97
	*/
	context.DrawImage(lisa, 183, 97)
}