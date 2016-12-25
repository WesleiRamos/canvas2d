package main

import "canvas2d"

var FlappyGame FlappyBirdGo

func main() {
	FlappyGame = FlappyBirdGo{}

	FlappyGame.canvas = canvas2d.NewCanvas(400, 500, "FlappyBirdGo")
	FlappyGame.context = FlappyGame.canvas.GetContext()

	FlappyGame.canvas.SetResizable(false)
	FlappyGame.canvas.SetFullScreen(true)
	FlappyGame.canvas.SetIcon("./data/icone.png")
	FlappyGame.canvas.SetLoadResources(FlappyGame.init)
	FlappyGame.canvas.SetLoopFunc(FlappyGame.Loop)

	FlappyGame.canvas.OnMouseDown(mousedown)

	FlappyGame.canvas.Show()
}

func mousedown(x, y float32, btn, mod int32) {
	if FlappyGame.gameStart {
		FlappyGame.SetPulando()
	} else {
		wd := float32(FlappyGame.canvas.Width) / 2
		hd := float32(FlappyGame.canvas.Height) / 2

		if (x > wd-50 && x < wd+50) && (y > hd+45 && y < hd+95) {
			FlappyGame.gameStart = true
		}
	}
}