package main

/*
go get -u github.com/faiface/beep
go get -u github.com/hajimehoshi/go-mp3
go get -u github.com/hajimehoshi/oto
go get -u github.com/pkg/errors
*/

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// sound source: http://soundbible.com/2197-Analog-Watch-Alarm.html

// Play alarm sound
func Play() {
	f, err := os.Open("alarm.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(4900 * time.Millisecond)
}
