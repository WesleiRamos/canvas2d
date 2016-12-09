package canvas2d

import "github.com/go-gl/gl/v2.1/gl"
import "github.com/go-gl/glfw/v3.2/glfw"

type KEY_PRESS func(key int32, mod int32)
type MOUSE_PRESS func(x, y float32, btn, mod int32)
type MOUSE_MOVE func(x, y float32)

type Canvas struct {
	Width, Height int
	Title         string
	LoopFunc      func()
	FullScreen    bool
	Resizable     bool
	Window        *glfw.Window
	KeyDown       KEY_PRESS
	KeyUp         KEY_PRESS
	KeyPress      KEY_PRESS
	MouseDown     MOUSE_PRESS
	MouseUp       MOUSE_PRESS
	MouseMove     MOUSE_MOVE
	KEYS          map[string]bool
}

func NewCanvas(w, h int, title string) Canvas {
	if w < 0 || h < 0 {
		panic("Width and Height must be positive")
	}
	return Canvas{Width: w, Height: h, Title: title, Resizable: true, KEYS: map[string]bool{}}
}

/*
	Funções "set"
*/

func (self *Canvas) SetLoopFunc(x func()) {
	/*
		Seta a função de loop do jogo
	*/
	self.LoopFunc = x
}

func (self *Canvas) SetFullScreen(full bool) {
	/*
		Seta o modo do display
		**Se a tela não for full screen não precisa chamar
	*/
	self.FullScreen = full
}

func (self *Canvas) SetResizable(r bool) {
	/*
		Seta se a tela poderá ser redimensionada ou não
	*/
	self.Resizable = r
}

func (self *Canvas) SetWatchKeys(keys string) {
	/*
		Cria chaves para as teclas

		exemplo:
			(camvas).SetWatchKeys("WASD")
	*/
	for key := range keys {
		self.KEYS[string(key)] = false
	}
}

func (self *Canvas) SetWatchKeyBool(k string, b bool) {
	/*
		Muda o valor das chaves teclas

		exemplo:
			(canvas).SetWatchKeyBool("A", true)
	*/
	if len(k) == 1 {
		self.KEYS[k] = b
	} else {
		panic("(canvas).SetWatchKeyBool(char string), string of len 1")
	}
}

/*
	Funções "get"
*/

func (self Canvas) GetContext() Context {
	return Context{fill{}, stroke{}}
}

func (self Canvas) GetKeyFromChar(char string) int32 {
	/*
		Retorna o numero da tecla correspondente ao caractere
	*/
	if len(char) == 1 {
		return int32([]rune(char)[0])
	}
	panic("(canvas).GetKeyFromChar(char string), string of len 1")
}

func (self Canvas) GetStringFromKeyCode(key int32) string {
	/*
		Retorna o caractere correspondente ao numero da tecla
	*/
	return string(key)
}

func (self *Canvas) GetWatchKey(key string) bool {
	/*
		Retorna o estado da tecla assistida

		exemplo:
			(canvas).GetWatchKey("A")
	*/
	if len(key) == 1 {
		return self.KEYS[key]
	}
	panic("(canvas).GetKeyWatch(key string), string of len 1")
}

/*
	Funções "on"
*/

func (self *Canvas) OnKeyDown(x KEY_PRESS) {
	// Define a ação de clicar no botão
	self.KeyDown = x
}

func (self *Canvas) OnKeyUp(x KEY_PRESS) {
	// Define a ação de doltar o botão
	self.KeyUp = x
}

func (self *Canvas) OnKeyPress(x KEY_PRESS) {
	// Define a ação de segurar o botão
	self.KeyPress = x
}

func (self *Canvas) OnMouseDown(x MOUSE_PRESS) {
	// Define a ação de clicar os botões do mouse
	self.MouseDown = x
}

func (self *Canvas) OnMouseUp(x MOUSE_PRESS) {
	// Define a ação de soltar os botões do mouse
	self.MouseUp = x
}

func (self *Canvas) OnMouseMove(x MOUSE_MOVE) {
	// Define a ação de mover o mouse
	self.MouseMove = x
}

/*
	Outras funções
*/

func (self Canvas) ApplicationExit() {
	// Fecha o programa
	self.Window.SetShouldClose(true)
}

func (self *Canvas) Show() {
	/*
		Inicia a janela
	*/

	// Inicializa o GLFW
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Samples, 4)

	if !self.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	/* Deixa em modo FullScreen */
	if self.FullScreen {
		self.setFullScreenMode()
	}

	/* Cria a janela */
	window, err := glfw.CreateWindow(self.Width, self.Height, self.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	/* Inicializa o OpenGL */
	if err := gl.Init(); err != nil {
		panic(err)
	}

	self.Window = window

	/* Callback events */

	// Quando a janela for redimensionada
	self.Window.SetSizeCallback(self.sizeCallback)

	// Quando alguma tecla for pressionada
	self.Window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

		if action == glfw.Press && self.KeyDown != nil {
			/* Quando a tecla for pressionada */
			self.KeyDown(int32(key), int32(mods))

		} else if action == glfw.Release && self.KeyUp != nil {
			/* Quando a tecla for solta */
			self.KeyUp(int32(key), int32(mods))

		} else if action == glfw.Repeat && self.KeyPress != nil {
			/* Se a tecla estiver ainda sendo pressionada */
			self.KeyPress(int32(key), int32(mods))
		}
	})

	// Quando algum botão do mouse for pressionado
	self.Window.SetMouseButtonCallback(func(w *glfw.Window, btn glfw.MouseButton, atc glfw.Action, mod glfw.ModifierKey) {
		posx, posy := w.GetCursorPos()

		if atc == glfw.Press && self.MouseDown != nil {
			// Quando o botão do mouse for pressionado
			self.MouseDown(float32(posx), float32(posy), int32(btn), int32(mod))

		} else if atc == glfw.Release && self.MouseUp != nil {
			// Quando o botão do mouse for solto
			self.MouseUp(float32(posx), float32(posy), int32(btn), int32(mod))
		}
	})

	// Pega a posição x e y do mouse quando for movimentado
	self.Window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		if self.MouseMove != nil {
			self.MouseMove(float32(x), float32(y))
		}
	})

	/*****************/

	if self.LoopFunc != nil {
		self.set2d()

		for !self.Window.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT)

			self.LoopFunc()

			gl.Flush()

			self.Window.SwapBuffers()
			glfw.PollEvents()
		}

	}
}

/* Funções privadas */

func (self Canvas) set2d() {
	/*
		Como o nome já diz, faz com que a projeção seja 2d
	*/
	gl.Viewport(0, 0, int32(self.Width), int32(self.Height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(self.Width), 0, float64(self.Height), -1, 1)
	gl.Scalef(1, -1, 1)
	gl.Translatef(0, -float32(self.Height), 0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func (self *Canvas) sizeCallback(_window *glfw.Window, w int, h int) {
	/*
		Quando a janela for redimensionada, altera também a projeção da janela
	*/
	self.Width = w
	self.Height = h
	self.set2d()
}

func (self *Canvas) setFullScreenMode() {
	/*
		Faz com que a janela entre em fullscreen
	*/
	monitor := glfw.GetPrimaryMonitor().GetVideoMode()
	self.Width = monitor.Width
	self.Height = monitor.Height
	glfw.WindowHint(glfw.Decorated, glfw.False)
}