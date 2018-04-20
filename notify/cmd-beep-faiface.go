package main

import (
	"os"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	f, _ := os.Open("/usr/share/sounds/ubuntu/ringtones/Harmonics.ogg")
	s, format, _ := wav.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(s)
}
