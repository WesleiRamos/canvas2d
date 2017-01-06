package main

import "fmt"
import "github.com/WesleiRamos/canvas2d"

type Obstaculo struct {
	x, y, cy, by, bh float32
	contouponto      bool
}

type FlappyBirdGo struct {
	canvas  canvas2d.Canvas
	context canvas2d.Context
	clock   canvas2d.Clock

	pontos         int32
	imagens        map[string]*canvas2d.Image
	fonte          *canvas2d.Font
	fontemenor     *canvas2d.Font
	obstaculos     []Obstaculo
	myposX, myposY float32
	pulando        bool
	pulocount      int32
	pulotempo      int32
	obscount       int32
	obsinterval    int32

	passaroAtual string
	countAnim    int32
	rotacao      float32
	caindo       bool

	gameOver  bool
	gameStart bool

	chao1, chao2 float32
}

func (self *FlappyBirdGo) init() {
	/*
		Inicializa as propriedades do jogo
	*/
	self.clock = canvas2d.NewClock()

	self.fonte = canvas2d.LoadFont("./data/04b_19.ttf", 40)
	self.fontemenor = canvas2d.LoadFont("./data/04b_19.ttf", 15)

	self.imagens = map[string]*canvas2d.Image{}
	self.CarregarImagens()

	self.ResetaPropriedades()
}

func (self *FlappyBirdGo) ResetaPropriedades() {
	/*
		Reseta propriedades
	*/
	self.obstaculos = []Obstaculo{}

	self.myposX = 40
	self.myposY = 100
	self.obsinterval = 100
	self.pulotempo = 20

	self.rotacao = 0
	self.pontos = 0

	self.chao1 = 0
	self.chao2 = 1024

	self.gameOver = false

	self.passaroAtual = "passaro1"
}

func (self *FlappyBirdGo) CarregarImagens() {
	/*
		Carrega as imagens do jogo
	*/
	imgs := []string{"canobaixo", "canocima", "fundo", "gameover", "getready", "passaro1",
		"passaro2", "passaro3", "passaro4", "play", "solo"}

	for img := range imgs {
		i := imgs[img]
		self.imagens[i] = canvas2d.LoadImage(fmt.Sprintf("./data/%s.png", i))
	}
}

func (self *FlappyBirdGo) MoveChao() {
	delta := float32(self.clock.DeltaTime()) * 15

	self.chao1 -= 10 * delta
	self.chao2 -= 10 * delta

	if self.chao1 <= -1050 {
		self.chao1 = 1024 + self.chao2

	} else if self.chao2 <= -1050 {
		self.chao2 = 1024 + self.chao1
	}
}

func (self *FlappyBirdGo) Gravidade() {
	delta := float32(self.clock.DeltaTime()) * 10
	/*
		Adiciona a gravidade (ta uma bosta precisa da animação)
	*/
	if self.pulando {

		if self.myposY > -40 {
			self.myposY -= 15 * delta
		}

		if self.rotacao > -20 {
			self.rotacao -= 8
		}

		self.pulocount++

		if self.pulocount == self.pulotempo {
			self.pulando = false
		}

	} else {
		self.myposY += 25 * delta

		if self.rotacao < 80 {
			self.rotacao += 5
			if self.rotacao < 80 {
				self.caindo = true
				self.passaroAtual = "passaro1"
			}
		} else {
			self.myposY += 5 * delta
		}

		if self.myposY+30 >= float32(self.canvas.Height)-70 {
			self.passaroAtual = "passaro4"
			self.gameOver = true
			self.myposY = (float32(self.canvas.Height) - 70) - 30
		}
	}
}

func (self *FlappyBirdGo) GerarObstaculos() {
	/*
		Gera obstaculos
	*/
	if self.obscount == self.obsinterval {

		// Pega a tela e remove o tamanho do chão
		sw := float32(self.canvas.Height) - 70
		py := float32(canvas2d.Random(0, int(sw)))

		// Evita que os canos fiquem pra cima ou pra baixo da tela
		if py < 60 {
			py = 60
		} else if py > (sw - 210) {
			py = sw - 210
		}

		cy := py - 301
		by := py + 150

		var bh float32 = 301

		if by < float32(self.canvas.Height)-301 {
			bh = float32(self.canvas.Height) - by
		}

		self.obstaculos = append(self.obstaculos, Obstaculo{x: float32(self.canvas.Width) + 20, y: float32(py), cy: cy, by: by, bh: bh})
		self.obscount = 0
	} else {
		self.obscount++
	}
}

func (self *FlappyBirdGo) UpdatePosObstaculos() {
	/*
		Atualiza a posição dos obstaculos
	*/
	obsRemove := []int{}
	for obs := range self.obstaculos {
		self.obstaculos[obs].x -= 10 * (float32(self.clock.DeltaTime()) * 15)

		if (self.obstaculos[obs].x + 75) < 0 {
			// Se passou da tela adiciona à lista de remoção
			obsRemove = append(obsRemove, obs)
		}
	}

	for rem := range obsRemove {
		// Se tiver algum elemento na lista remove
		self.RemoveObstaculos(obsRemove[rem])
	}
}

func (self *FlappyBirdGo) RemoveObstaculos(pos int) {
	self.obstaculos = append(self.obstaculos[:pos], self.obstaculos[pos+1:]...)
}

func (self *FlappyBirdGo) DrawObjects() {
	/*
		Desenha os elementos do jogo
		Cenário
		Canos
		Passaro
		Chão
	*/
	self.context.DrawImage(self.imagens["fundo"], 0, 0, float32(self.canvas.Width), float32(self.canvas.Height))

	for o := range self.obstaculos {
		obs := self.obstaculos[o]

		self.context.DrawImage(self.imagens["canocima"], obs.x, obs.cy, 70, 301)
		self.context.DrawImage(self.imagens["canobaixo"], obs.x, obs.by, 70, obs.bh)
	}

	self.context.DrawImage(self.imagens["solo"], self.chao1, float32(self.canvas.Height)-70)
	self.context.DrawImage(self.imagens["solo"], self.chao2, float32(self.canvas.Height)-70)

	if self.gameStart {
		self.DrawPassaro()
		self.DrawPontos()
	} else {
		self.context.DrawImage(self.imagens["getready"], float32(self.canvas.Width)/2-173, float32(self.canvas.Height/2)-60)
		self.context.DrawImage(self.imagens["play"], float32(self.canvas.Width)/2-50, float32(self.canvas.Height/2)+45, 100, 50)
	}

	if self.gameOver {
		self.context.DrawImage(self.imagens["gameover"], float32(self.canvas.Width)/2-173, float32(self.canvas.Height/2)-38)
		self.context.DrawImage(self.imagens["play"], float32(self.canvas.Width)/2-50, float32(self.canvas.Height/2)+45, 100, 50)
	}
}

func (self *FlappyBirdGo) DrawPassaro() {
	// 22.5 = (45/2) : 45 largura do passaro
	// 15   = (30/2) : 30 altura do passaro
	self.context.Translate(self.myposX+22.5, self.myposY+15)
	self.context.Rotate(self.rotacao)

	self.context.DrawImage(self.imagens[self.passaroAtual], -22.5, -15, 45, 30)

	self.context.Restore()
}

func (self *FlappyBirdGo) DrawPontos() {
	/*
		Mostra os pontos
	*/
	self.context.Fill.Font = self.fonte

	p := float32(self.canvas.Width) / 2
	c := fmt.Sprintf("%d", self.pontos)
	d := p - (10 * float32(len(c)))

	self.context.Fill.Style = canvas2d.NewColor(0, 0, 0)
	self.context.Fill.Text(c, d-2, 18)

	self.context.Fill.Style = canvas2d.NewColor(255, 255, 255)
	self.context.Fill.Text(c, d, 20)

	self.context.Fill.Font = self.fontemenor
	self.context.Fill.Text(fmt.Sprintf("FPS: %d", self.clock.FPS()), 20, 10)
}

func (self *FlappyBirdGo) ChecaGameOver() {
	for o := range self.obstaculos {
		obs := self.obstaculos[o]

		// 45 é a largura do passaro
		// 70 é a largura do cano
		if self.myposX+45 >= obs.x && self.myposX <= obs.x+70 {
			if self.myposY < obs.y || (self.myposY+30) > obs.by {
				// Colidiu cima ou baixo
				self.gameOver = true
				self.passaroAtual = "passaro4"
			}
		} else if self.myposX > obs.x+70 && !obs.contouponto {
			self.pontos++
			self.obstaculos[o].contouponto = true
		}
	}
}

func (self *FlappyBirdGo) Animacao() {
	/*
		Animação
	*/
	if !self.caindo {
		if self.countAnim == 5 {
			self.passaroAtual = "passaro2"
		} else if self.countAnim == 10 {
			self.passaroAtual = "passaro1"
		} else if self.countAnim == 15 {
			self.passaroAtual = "passaro3"
		} else if self.countAnim == 20 {
			self.passaroAtual = "passaro1"
			self.countAnim = 0
		}
		self.countAnim++
	}
}

func (self *FlappyBirdGo) Loop() {
	/*
		Loop do jogo
	*/
	self.clock.Tick()

	if self.gameStart {
		if !self.gameOver {
			self.ChecaGameOver()
			self.GerarObstaculos()
			self.UpdatePosObstaculos()
			self.Animacao()
		}
		self.Gravidade()
	}
	if !self.gameOver {
		self.MoveChao()
	}
	self.DrawObjects()
}

func (self *FlappyBirdGo) SetPulando() {
	/*
		reseta o pulo
	*/
	if !self.gameOver {
		self.myposY -= 1
		self.pulocount = 0
		self.pulando = true
		self.caindo = false
	}
}
