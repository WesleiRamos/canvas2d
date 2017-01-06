package canvas2d

import "time"

type Clock struct {
	fps         int32
	passou      float64
	frames      int32
	deltaTime   float64
	ultimoTempo time.Time
}

func (self *Clock) Tick() {
	self.frames++
	agora := time.Now()
	self.deltaTime = agora.Sub(self.ultimoTempo).Seconds()
	self.passou += self.deltaTime
	self.ultimoTempo = agora

	if self.passou >= 1 {
		self.fps = self.frames
		self.frames = 0
		self.passou = 0
	}
}

func (self *Clock) FPS() int32 {
	return self.fps
}

func (self *Clock) DeltaTime() float64 {
	return self.deltaTime
}

func NewClock() Clock {
	clock := Clock{}
	clock.Tick()
	return clock
}