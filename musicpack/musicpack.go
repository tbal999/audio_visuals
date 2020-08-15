package musicpack

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//Not used in the application. But you can add music to it using this.
var (
	streamer   beep.StreamSeekCloser
	format     beep.Format
	stop       = make(chan bool)
	down       = make(chan bool)
	up         = make(chan bool)
	quiet      = make(chan bool)
	loud       = make(chan bool)
	soundconst = 0
)

func Play(music string) {
	f, err := os.Open(music)
	if err != nil {
		return
	}
	streamer, format, err = mp3.Decode(f)
	if err != nil {
		return
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	sample := beep.ResampleRatio(4, 1, volume)

	speaker.Play(sample)
	for {
		select {
		case <-stop:
			speaker.Close()
		case <-down:
			if soundconst > -4 {
				speaker.Lock()
				sample.SetRatio(sample.Ratio() - 0.1)
				speaker.Unlock()
				soundconst--
			}
		case <-up:
			if soundconst < 4 {
				speaker.Lock()
				sample.SetRatio(sample.Ratio() + 0.1)
				speaker.Unlock()
				soundconst++
			}
		case <-quiet:
			speaker.Lock()
			volume.Volume--
			speaker.Unlock()
		case <-loud:
			speaker.Lock()
			volume.Volume++
			speaker.Unlock()
		}
	}
}

func Stop() {
	stop <- true
}

func Down() {
	down <- true
}

func Up() {
	up <- true
}

func Quiet() {
	quiet <- true
}

func Loud() {
	loud <- true
}
