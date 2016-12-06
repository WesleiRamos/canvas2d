package canvas2d

import "github.com/go-gl/gl/v2.1/gl"
import "github.com/go-gl/mathgl/mgl32"

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

/* Desenha circulos */
func (self fill) Arc(x, y, raio float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	arc := mgl32.Circle(raio, raio, 90)

	gl.Begin(gl.POLYGON)

	for point := range arc {
		px, py := arc[point].Elem()
		gl.Vertex2f(px+x, py+y)
	}

	gl.End()
}

func (self stroke) Arc(x, y, raio float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)

	arc := mgl32.Circle(raio, raio, 90)

	gl.Begin(gl.LINE_LOOP)

	for point := range arc {
		px, py := arc[point].Elem()
		if px != 0 && py != 0 {
			gl.Vertex2f(px+x, py+y)
		}
	}

	gl.End()
}