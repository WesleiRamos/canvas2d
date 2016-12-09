package canvas2d

import "math"
import "github.com/go-gl/gl/v2.1/gl"

const twoPi float64 = 2.0 * math.Pi

type Color struct {
	R, G, B float32
}

/* Desenha e preenche */
type fill struct {
	Style Color
}

/* Desenha somente o contorno */
type stroke struct {
	Style Color
}

type Context struct {
	Fill   fill
	Stroke stroke
}

/* Cor de fundo */
func (self Context) Background(cor Color) {
	gl.ClearColor(cor.R, cor.G, cor.B, 1.0)
}

/* Desenha retangulo com preenchimento */
func (self fill) Rect(x, y, w, h float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	gl.Begin(gl.QUADS)
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+w, y)
	gl.Vertex2f(x+w, y+h)
	gl.Vertex2f(x, y+h)
	gl.End()
}

/* Desenha retangulo com somente o contorno */
func (self stroke) Rect(x, y, w, h float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+w, y)
	gl.Vertex2f(x+w, y+h)
	gl.Vertex2f(x, y+h)
	gl.End()
}

/* Desenha linha */
func (self stroke) Line(x1, y1, x2, y2 float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	gl.Begin(gl.LINES)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()
}

/* Desenha circulo preenchido */
func (self fill) Circle(x, y, raio float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	triangles := float64(int(raio))

	if triangles > 90 {
		triangles = 90
	} else if triangles <= 15 {
		triangles = float64(int((triangles / 180) * 360))
	}

	var i float64 = 0
	twopi := twoPi / triangles

	gl.Begin(gl.TRIANGLE_FAN)
	gl.Vertex2f(x, y)

	for ; i <= triangles; i++ {
		_x := x + (raio * float32(math.Cos(i*twopi)))
		_y := y + (raio * float32(math.Sin(i*twopi)))
		gl.Vertex2f(_x, _y)
	}
	gl.End()
}

/* Desenha circulo com somente o contorno */
func (self stroke) Circle(x, y, raio float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	points := float64(int(raio))

	if points > 90 {
		points = 90
	} else if points <= 15 {
		points = float64(int((points / 180) * 360))
	}

	var i float64 = 0
	twopi := twoPi / points

	gl.Begin(gl.LINE_LOOP)

	for ; i <= points; i++ {
		_x := x + (raio * float32(math.Cos(i*twopi)))
		_y := y + (raio * float32(math.Sin(i*twopi)))
		gl.Vertex2f(_x, _y)
	}
	gl.End()
}