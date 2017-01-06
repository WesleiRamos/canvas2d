package canvas2d

const mult float32 = 0.00392156862745098

type Color struct {
	R, G, B, A float32
}

func (self Color) RGBA() (r, g, b, a float32) {
	r = self.R * 255
	g = self.G * 255
	b = self.B * 255
	a = self.A

	return
}

func NewColor(r, g, b float32, a ...float32) *Color {
	var alpha float32 = 1

	if len(a) > 0 {
		alpha = a[0]

		if alpha > 1 {
			alpha = 1
		}
	}

	return &Color{r * mult, g * mult, b * mult, alpha}
}