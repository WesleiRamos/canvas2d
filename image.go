package canvas2d

import "os"
import "image"
import "image/draw"
import _ "image/png"
import _ "image/jpeg"

type Image struct {
	data      *image.NRGBA
	width     int32
	height    int32
	imgNumber uint32
}

func (self Image) GetSize() (int32, int32) {
	return self.width, self.height
}

func (self Image) GetNumber() uint32 {
	/* Numero da textura */
	return self.imgNumber
}

func (self *Image) SetImgNumber(number uint32) {
	self.imgNumber = number
}

func loadImage(filepath string) *Image {
	file, err := os.Open(filepath)
	if err != nil {
		panic("(Load Image) Image not found")
	}
	defer file.Close()

	imagem, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	nrgba := image.NewNRGBA(imagem.Bounds())
	draw.Draw(nrgba, nrgba.Bounds(), imagem, image.Point{0, 0}, draw.Src)

	return &Image{data: nrgba, width: int32(nrgba.Rect.Size().X), height: int32(nrgba.Rect.Size().Y)}
}

func loadIcon(filepath string) []image.Image {
	file, err := os.Open(filepath)
	if err != nil {
		panic("(Load Icon) Icon not found")
	}
	defer file.Close()

	icon, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return []image.Image{icon}
}