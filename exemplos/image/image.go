package main

import "canvas2d"

var context canvas2d.Context
var lisa canvas2d.Image

func main() {
	canvas := canvas2d.NewCanvas(600, 400, "IMAGEM")
	context = canvas.GetContext()

	canvas.SetResizable(false)
	canvas.SetLoadResources(loadResources)
	canvas.SetLoopFunc(Loop)

	canvas.Show()
}

func loadResources() {
	context.Background(canvas2d.Color{1, 1, 1})
	lisa = canvas2d.LoadImage("lisa.jpg")
}

func Loop() {
	context.DrawImage(lisa, 183, 97)
}