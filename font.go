package canvas2d

import "os"
import "github.com/go-gl/gltext"

type Font struct {
	font *gltext.Font
}

func LoadFont(file string, size int32) *Font {
	font, err := os.Open(file)
	if err != nil {
		panic("(Load Font) Font not found")
	}
	f, err := gltext.LoadTruetype(font, size, 32, 255, gltext.LeftToRight)
	if err != nil {
		panic(err)
	}
	return &Font{f}
}