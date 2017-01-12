package canvas2d

import "math"
import "github.com/go-gl/gl/v2.1/gl"

const twoPi float64 = 2.0 * math.Pi

/* Desenha e preenche */
type fill struct {
	Style *Color
	Font  *Font
}

/* Desenha somente o contorno */
type stroke struct {
	Style *Color
}

type Context struct {
	Fill   fill
	Stroke stroke
}

/* Cor de fundo */
func (self Context) Background(cor *Color) {
	gl.ClearColor(cor.R, cor.G, cor.B, 1.0)
}

func (self Context) Translate(x, y float32) {
	/*
		Move para
	*/
	gl.Translatef(x, y, 0)
}

func (self Context) Rotate(angle float32) {
	/*
		Restaura
	*/
	gl.Rotatef(angle, 0, 0, 1)
}

func (self Context) Restore() {
	/*
		Restaura a matrix
	*/
	gl.LoadIdentity()
}

/* Desenha retangulo com preenchimento */
func (self fill) Rect(x, y, w, h float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	gl.Begin(gl.QUADS)
	{
		gl.Vertex2f(x, y)
		gl.Vertex2f(x+w, y)
		gl.Vertex2f(x+w, y+h)
		gl.Vertex2f(x, y+h)
	}
	gl.End()
}

/* Desenha retangulo com somente o contorno */
func (self stroke) Rect(x, y, w, h float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	gl.Begin(gl.LINE_LOOP)
	{
		gl.Vertex2f(x, y)
		gl.Vertex2f(x+w, y)
		gl.Vertex2f(x+w, y+h)
		gl.Vertex2f(x, y+h)
	}
	gl.End()
}

/* Desenha linha */
func (self stroke) Line(x1, y1, x2, y2 float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	gl.Begin(gl.LINES)
	{
		gl.Vertex2f(x1, y1)
		gl.Vertex2f(x2, y2)
	}
	gl.End()
}

// Lines
func (self stroke) Lines(p *[][]float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	jb := *p

	gl.Begin(gl.LINE_LOOP)
	{
		for i := range jb {
			b := jb[i]
			gl.Vertex2f(b[0], b[1])
		}
	}
	gl.End()
}

/* Desenha circulo preenchido */
func (self fill) Circle(x, y, raio float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	triangles := float64(int(raio))

	if triangles > 90 {
		triangles = 90
	} else if triangles <= 15 {
		triangles = float64(int((triangles / 180) * 360))
	}

	var i float64 = 0
	twopi := twoPi / triangles

	gl.Begin(gl.TRIANGLE_FAN)
	{
		gl.Vertex2f(x, y)

		for ; i <= triangles; i++ {
			_x := raio * float32(math.Cos(i*twopi))
			_y := raio * float32(math.Sin(i*twopi))
			gl.Vertex2f(x+_x, y+_y)
		}
	}
	gl.End()
}

/* Desenha circulo com somente o contorno */
func (self stroke) Circle(x, y, raio float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	points := float64(int(raio))

	if points > 90 {
		points = 90
	} else if points <= 15 {
		points = float64(int((points / 180) * 360))
	}

	var i float64 = 0
	twopi := twoPi / points

	gl.Begin(gl.LINE_LOOP)
	{
		for ; i <= points; i++ {
			_x := raio * float32(math.Cos(i*twopi))
			_y := raio * float32(math.Sin(i*twopi))
			gl.Vertex2f(x+_x, y+_y)
		}
	}
	gl.End()
}

// Poligono
func (self fill) Polygon(x, y float32, v *[][]float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)

	vertices := *v

	gl.Begin(gl.POLYGON)
	{
		for i := range vertices {
			b := vertices[i]
			gl.Vertex2f(x+b[0], y+b[1])
		}
	}
	gl.End()
}

/* Escreve texto */
func (self fill) Text(text string, x, y float32) {
	gl.Color4f(self.Style.R, self.Style.G, self.Style.B, self.Style.A)
	self.Font.font.Printf(x, y, text)
}

/* Desenha a imagem (na verdade é uma textura) */
func (self *Context) DrawImage(imagem *Image, p ...float32) {
	if (len(p) > 0) && (len(p)%2 == 0) && (len(p) <= 8) {

		gl.Color3f(1.0, 1.0, 1.0)
		gl.BindTexture(gl.TEXTURE_2D, imagem.imgNumber)

		x, y := p[0], p[1]
		w, h := float32(imagem.width), float32(imagem.height)

		var x1, x2, y1, y2 float32 = 0, 1, 0, 1

		switch len(p) {
		case 4:
			w = p[2]
			h = p[3]
		case 8:
			w = p[2]
			h = p[3]
			x1 = p[4]
			x2 = p[6]
			y1 = p[5]
			y2 = p[7]
		case 6:
			x1 = p[2]
			x2 = p[4]
			y1 = p[3]
			y2 = p[5]
		}

		gl.Begin(gl.QUADS)
		{
			gl.TexCoord2f(x1, y1)
			gl.Vertex2f(x, y)
			gl.TexCoord2f(x2, y1)
			gl.Vertex2f(x+w, y)
			gl.TexCoord2f(x2, y2)
			gl.Vertex2f(x+w, y+h)
			gl.TexCoord2f(x1, y2)
			gl.Vertex2f(x, y+h)
		}
		gl.End()

	} else {
		panic("(Draw image) Parameters error, number of parameters : " + string(len(p)))
	}
}

func DestroyImage(image Image) {
	n := image.GetNumber()
	gl.DeleteTextures(1, &n)
}

func LoadImage(file string) *Image {
	image := loadImage(file)

	var imgNum uint32

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &imgNum)
	gl.BindTexture(gl.TEXTURE_2D, imgNum)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, image.width, image.height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image.data.Pix))

	image.SetImgNumber(imgNum)

	return image
}