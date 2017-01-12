package canvas2d

import "bitbucket.org/StephenPatrick/go-winaudio/winaudio"

var winaudioInicializado bool

type Audio struct {
	audio   winaudio.Audio
	playing bool
}

func (self *Audio) Play() {
	self.playing = true
	self.audio.Play()
}

func (self *Audio) Pause() {
	self.playing = false
	self.audio.Pause()
}

func (self *Audio) Stop() {
	self.playing = false
	self.audio.Stop()
}

func (self *Audio) IsPlaying() bool {
	return self.playing
}

func NewAudio(filepath string) *Audio {
	// Windows Only
	if !winaudioInicializado {
		winaudio.InitWinAudio()
		winaudioInicializado = true
	}
	wav, err := winaudio.LoadWav(filepath)
	if err != nil {
		panic(err)
	}
	return &Audio{audio: wav}
}