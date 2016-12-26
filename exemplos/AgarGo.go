package main

import "github.com/WesleiRamos/canvas2d"
import "math"

type Cell struct {
	x, y, r float32
	anim    int
	color   canvas2d.Color
}

type AgarGo struct {
	cells                    []Cell
	canvas                   canvas2d.Canvas
	context                  canvas2d.Context
	gameInitialized          bool
	myPosX, myPosY, myRadius float32
	updateX, updateY         float32
	addVel                   float32
}

func main() {
	game := AgarGo{cells: []Cell{}}
	game.canvas = canvas2d.NewCanvas(600, 400, "AgarGo")
	game.context = game.canvas.GetContext()

	//game.canvas.SetFullScreen(true)
	game.canvas.SetResizable(false)
	game.canvas.SetLoopFunc(game.loop)
	game.canvas.SetLoadResources(game.init)

	game.canvas.OnKeyDown(game.keyDown)
	game.canvas.OnMouseDown(game.mouseDown)

	game.canvas.Show()
}

func (self *AgarGo) init() {
	self.context.Background(canvas2d.Color{1.0, 1.0, 1.0})

	celln := 100

	for i := 0; i < celln; i++ {
		x := canvas2d.Random(0, self.canvas.Width)
		y := canvas2d.Random(0, self.canvas.Height)
		r := float32(canvas2d.Random(0, 100)) / 100
		g := float32(canvas2d.Random(0, 100)) / 100
		b := float32(canvas2d.Random(0, 100)) / 100
		self.cells = append(self.cells, Cell{float32(x), float32(y), 10, 1, canvas2d.Color{r, g, b}})
	}

	self.myPosX = float32(self.canvas.Width / 2)
	self.myPosY = float32(self.canvas.Height / 2)
	self.myRadius = 16
	self.addVel = 1.5
}

func (self *AgarGo) loop() {
	self.updatePos()
	self.detectCollision()
	self.draw()
}

func (self *AgarGo) updatePos() {
	self.myPosX += self.updateX
	self.myPosY += self.updateY
}

func (self *AgarGo) mouseDown(x, y float32, btn, mod int32) {
	self.updateX = x - self.myPosX
	self.updateY = y - self.myPosY
	dist := float32(math.Sqrt(math.Pow(float64(self.updateX), 2) + math.Pow(float64(self.updateY), 2)))
	self.updateX = self.updateX / dist
	self.updateY = self.updateY / dist

	self.updateX *= self.addVel
	self.updateY *= self.addVel
}

func (self *AgarGo) keyDown(key, mod int32) {
	if key == 67 {
		// KEY "C"
		self.canvas.ApplicationExit()
	}
}

func (self *AgarGo) draw() {
	//lines (grid)
	self.context.Stroke.Style = canvas2d.Color{0.8, 0.8, 0.8}

	xm := self.canvas.Width / 14
	ym := self.canvas.Height / 7

	for x := 1; x < 14; x++ {
		_x := float32(xm * x)
		self.context.Stroke.Line(_x, 0, _x, float32(self.canvas.Height))
	}

	for y := 1; y < 7; y++ {
		_y := float32(ym * y)
		self.context.Stroke.Line(0, _y, float32(self.canvas.Width), _y)
	}

	// cells
	for c := range self.cells {
		cell := self.cells[c]
		self.context.Fill.Style = cell.color
		self.context.Fill.Circle(cell.x, cell.y, cell.r)

		// Cell Animation
		if cell.anim == 1 {
			if cell.r >= 12 {
				self.cells[c].anim = 0
			}
			self.cells[c].r += 0.01
		} else {
			if cell.r <= 10 {
				self.cells[c].anim = 1
			}
			self.cells[c].r -= 0.01
		}
	}

	self.context.Fill.Style = canvas2d.Color{0.2, 0.2, 0.2}
	self.context.Fill.Circle(self.myPosX, self.myPosY, self.myRadius)
}

func (self *AgarGo) removeCell(pos int) {
	self.cells = append(self.cells[:pos], self.cells[pos+1:]...)
}

func (self *AgarGo) detectCollision() {
	cellRem := []int{}
	for cellnum := range self.cells {
		cell := self.cells[cellnum]

		if self.collision(cell.x, cell.y, cell.r) {
			cellRem = append(cellRem, cellnum)
		}
	}

	for cellnum := range cellRem {
		self.removeCell(cellRem[cellnum])
		self.myRadius += 2

		if self.myRadius < 70 {
			self.addVel = 2 - (self.myRadius / 40)
		}
	}
}

func (self *AgarGo) collision(x, y, r float32) bool {
	// circle colision
	c1 := float64(self.myPosX - x)
	c2 := float64(self.myPosY - y)
	return math.Sqrt(c1*c1+c2*c2) < float64((self.myRadius-5)+r)
}