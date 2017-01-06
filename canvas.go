package canvas2d

import "image"
import "runtime"
import "github.com/go-gl/gl/v2.1/gl"
import "github.com/go-gl/glfw/v3.2/glfw"

type KEY_PRESS func(key int32, mod int32)
type MOUSE_PRESS func(x, y float32, btn, mod int32)
type MOUSE_MOVE func(x, y float32)
type WINDOW_RESIZE func(w, h int)

func init() {
	runtime.LockOSThread()
}

type Canvas struct {
	Width, Height int
	title         string
	loopFunc      func()
	fullScreen    bool
	resizable     bool
	loadResources func()
	swapInterval  int
	cursor        bool

	icon         []image.Image
	glfwInit     bool
	glInit       bool
	Window       *glfw.Window
	windowResize WINDOW_RESIZE
	keyDown      KEY_PRESS
	keyUp        KEY_PRESS
	keyPress     KEY_PRESS
	mouseDown    MOUSE_PRESS
	mouseUp      MOUSE_PRESS
	mouseMove    MOUSE_MOVE
}

func NewCanvas(w, h int, title string) Canvas {
	if w < 0 || h < 0 {
		panic("Width and Height need be positive")
	}
	return Canvas{Width: w, Height: h, title: title, resizable: true}
}

/*
	Funções "set"
*/

func (self *Canvas) SetLoadResources(x func()) {
	/*
		Seta uma função para carregar elementos ao jogo após iniciar o opengl
	*/
	self.loadResources = x
}

func (self *Canvas) SetLoopFunc(x func()) {
	/*
		Seta a função de loop do jogo
	*/
	self.loopFunc = x
}

func (self *Canvas) SetFullScreen(full bool) {
	/*
		Seta o modo do display
		**Se a tela não for full screen não precisa chamar
	*/
	self.fullScreen = full
}

func (self *Canvas) SetResizable(b bool) {
	/*
		Seta se a tela poderá ser redimensionada ou não
	*/
	self.resizable = b
}

func (self *Canvas) SetSwapInterval(i int) {
	self.swapInterval = i
}

func (self *Canvas) SetIcon(iconpath string) {
	self.icon = loadIcon(iconpath)
}

/*
	Funções "get"
*/

func (self Canvas) GetContext() Context {
	return Context{fill{}, stroke{}}
}

/*
	Funções "on"
*/

func (self *Canvas) OnKeyDown(x KEY_PRESS) {
	// Define a ação de clicar no botão
	self.keyDown = x
}

func (self *Canvas) OnKeyUp(x KEY_PRESS) {
	// Define a ação de doltar o botão
	self.keyUp = x
}

func (self *Canvas) OnKeyPress(x KEY_PRESS) {
	// Define a ação de segurar o botão
	self.keyPress = x
}

func (self *Canvas) OnMouseDown(x MOUSE_PRESS) {
	// Define a ação de clicar os botões do mouse
	self.mouseDown = x
}

func (self *Canvas) OnMouseUp(x MOUSE_PRESS) {
	// Define a ação de soltar os botões do mouse
	self.mouseUp = x
}

func (self *Canvas) OnMouseMove(x MOUSE_MOVE) {
	// Define a ação de mover o mouse
	self.mouseMove = x
}

func (self *Canvas) OnWindowResize(x WINDOW_RESIZE) {
	//	Seta a função que será chamada quando a janela for redimensionada
	self.windowResize = x
}

func (self *Canvas) EnableCursor() {
	/*
		Mostra o cursor do mouse
	*/
	self.cursor = true
	self.enableDisableCursor()
}

func (self *Canvas) DisableCursor() {
	/*
		Esconde o cursor do mouse
	*/
	self.cursor = false
	self.enableDisableCursor()
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

	self.glfwInit = true

	glfw.WindowHint(glfw.Samples, 4)

	if !self.resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	/* Cria a janela */
	window, err := self.newWindow()
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if self.swapInterval > 0 {
		glfw.SwapInterval(self.swapInterval)
	}

	/* Inicializa o OpenGL */
	if err := gl.Init(); err != nil {
		panic(err)
	}

	self.glInit = true

	if self.loadResources != nil {
		self.loadResources()
	}

	self.Window = window

	if len(self.icon) == 1 {
		self.Window.SetIcon(self.icon)
	}

	/* Callback events */

	// Quando a janela for redimensionada
	self.Window.SetSizeCallback(self.sizeCallback)

	// Quando alguma tecla for pressionada
	self.Window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

		if action == glfw.Press && self.keyDown != nil {
			/* Quando a tecla for pressionada */
			self.keyDown(int32(key), int32(mods))

		} else if action == glfw.Release && self.keyUp != nil {
			/* Quando a tecla for solta */
			self.keyUp(int32(key), int32(mods))

		} else if action == glfw.Repeat && self.keyPress != nil {
			/* Se a tecla estiver ainda sendo pressionada */
			self.keyPress(int32(key), int32(mods))
		}
	})

	// Quando algum botão do mouse for pressionado
	self.Window.SetMouseButtonCallback(func(w *glfw.Window, btn glfw.MouseButton, atc glfw.Action, mod glfw.ModifierKey) {
		posx, posy := w.GetCursorPos()

		if atc == glfw.Press && self.mouseDown != nil {
			// Quando o botão do mouse for pressionado
			self.mouseDown(float32(posx), float32(posy), int32(btn), int32(mod))

		} else if atc == glfw.Release && self.mouseUp != nil {
			// Quando o botão do mouse for solto
			self.mouseUp(float32(posx), float32(posy), int32(btn), int32(mod))
		}
	})

	// Pega a posição x e y do mouse quando for movimentado
	self.Window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		if self.mouseMove != nil {
			self.mouseMove(float32(x), float32(y))
		}
	})

	/*****************/

	if self.loopFunc != nil {
		self.set2d()

		for !self.Window.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT)

			self.loopFunc()

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
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LoadIdentity()
}

func (self *Canvas) sizeCallback(_window *glfw.Window, w, h int) {
	/*
		Quando a janela for redimensionada, altera também a projeção da janela
	*/
	self.Width = w
	self.Height = h
	self.set2d()

	if self.windowResize != nil {
		self.windowResize(w, h)
	}
}

func (self *Canvas) newWindow() (*glfw.Window, error) {
	/*
		Faz com que a janela entre em fullscreen
	*/
	if self.fullScreen {
		monitor := glfw.GetPrimaryMonitor().GetVideoMode()
		self.Width = monitor.Width
		self.Height = monitor.Height

		return glfw.CreateWindow(self.Width, self.Height, self.title, glfw.GetPrimaryMonitor(), nil)
	} else {
		return glfw.CreateWindow(self.Width, self.Height, self.title, nil, nil)
	}
}

func (self *Canvas) enableDisableCursor() {
	if self.glfwInit {
		cursor := glfw.CursorNormal
		if !self.cursor {
			cursor = glfw.CursorHidden
		}
		self.Window.SetInputMode(glfw.CursorMode, cursor)
	}
}
