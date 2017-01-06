package main

import "github.com/WesleiRamos/canvas2d"

var context canvas2d.Context
var font *canvas2d.Font

func main() {
	canvas := canvas2d.NewCanvas(600, 400, "LOREM IP... IP... IPSUM")
	context = canvas.GetContext()

	canvas.SetLoadResources(loadResources)
	canvas.SetLoopFunc(Loop)

	canvas.Show()
}

func loadResources() {
	context.Background(canvas2d.NewColor(255, 255, 255))
	font = canvas2d.LoadFont("../res/04b_19.ttf", 30)
	context.Fill.Font = font
	context.Fill.Style = canvas2d.NewColor(0, 0, 0)
}

func Loop() {
	context.Fill.Text("LOREM IP... IP... IPSUM", 50, 175)
}