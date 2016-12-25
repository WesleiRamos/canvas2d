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
	Font  Font
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

/*
	Move para
*/
func (self Context) Translate(x, y float32) {
	gl.Translatef(x, y, 0)
}

/*
	Restaura
*/
func (self Context) Rotate(angle float32) {
	gl.Rotatef(angle, 0, 0, 1)
}

/*
	Restaura a matrix
*/
func (self Context) Restore() {
	gl.LoadIdentity()
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
		_x := raio * float32(math.Cos(i*twopi))
		_y := raio * float32(math.Sin(i*twopi))
		gl.Vertex2f(x+_x, y+_y)
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
		_x := raio * float32(math.Cos(i*twopi))
		_y := raio * float32(math.Sin(i*twopi))
		gl.Vertex2f(x+_x, y+_y)
	}
	gl.End()
}

/* Escreve texto */
func (self fill) Text(text string, x, y float32) {
	gl.Color3f(self.Style.R, self.Style.G, self.Style.B)
	self.Font.font.Printf(x, y, text)
}

/* Desenha a imagem (na verdade é uma textura) */
func (self *Context) DrawImage(imagem Image, p ...float32) {
	if len(p) == 2 || len(p) == 4 {

		gl.Color3f(1.0, 1.0, 1.0)
		gl.BindTexture(gl.TEXTURE_2D, imagem.imgNumber)

		x, y := p[0], p[1]
		w, h := float32(imagem.width), float32(imagem.height)
		if len(p) == 4 {
			w = p[2]
			h = p[3]
		}

		gl.Begin(gl.QUADS)
		gl.TexCoord2i(0, 0)
		gl.Vertex2f(x, y)
		gl.TexCoord2i(1, 0)
		gl.Vertex2f(x+w, y)
		gl.TexCoord2i(1, 1)
		gl.Vertex2f(x+w, y+h)
		gl.TexCoord2i(0, 1)
		gl.Vertex2f(x, y+h)
		gl.End()

	} else {
		panic("(Draw image) Parameters error, number of parameters : " + string(len(p)))
	}
}

func DestroyImage(image Image) {
	n := image.GetNumber()
	gl.DeleteTextures(1, &n)
}

func LoadImage(file string) Image {
	image, rgba := loadImage(file)

	var imgNum uint32

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &imgNum)
	gl.BindTexture(gl.TEXTURE_2D, imgNum)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		image.width,
		image.height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	image.SetImgNumber(imgNum)

	return image
}