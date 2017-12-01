package reader

import (
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"time"

	golang_tts "github.com/leprosus/golang-tts"
)

type Reader struct {
	Reading bool
	Mutex   sync.Mutex
}

func NewReader() *Reader {
	return &Reader{
		Reading: false,
		Mutex:   sync.Mutex{},
	}
}

func (r *Reader) Read(value string) {
	read := false
	r.Mutex.Lock()
	if !r.Reading {
		read = true
		r.Reading = true
	}
	r.Mutex.Unlock()
	if read {
		go func() {
			Read(value)
			time.Sleep(10 * time.Second)
			r.Mutex.Lock()
			r.Reading = false
			r.Mutex.Unlock()
		}()
	}

}

func Read(value string) {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	polly := golang_tts.New(accessKey, secretKey)

	polly.Format(golang_tts.MP3)
	polly.Voice(golang_tts.Celine)

	bytes, err := polly.Speech(value)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./result.mp3", bytes, 0644)

	cmd := exec.Command("vlc", "-Idummy", "./result.mp3", "vlc://quit")
	cmd.Run()

}
