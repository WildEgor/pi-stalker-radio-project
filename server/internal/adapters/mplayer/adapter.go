package mplayer

import (
	"io"
	"os/exec"
	"runtime"
	"strings"
)

type IPlayer interface {
	Play(url string) error
	Mute()
	Pause()
	IncVolume()
	DecVolume()
	Stop()
	Close()
}

var _ IPlayer = (*MPlayer)(nil)

// MPlayer docs https://www.mankier.com/1/mplayer#Examples_(TL;DR)
type MPlayer struct {
	name    string
	command *exec.Cmd
	in      io.WriteCloser
	out     io.ReadCloser
	pipe    chan io.ReadCloser
	done    chan struct{}
}

func NewMPlayer() *MPlayer {
	name := getExecutable()

	return &MPlayer{
		name: name,
		pipe: make(chan io.ReadCloser),
		done: make(chan struct{}),
	}
}

func (p *MPlayer) Play(url string) error {
	var err error

	// check supported lists
	isPlaylist := strings.HasSuffix(url, ".m3u") ||
		strings.HasSuffix(url, ".pls")

	if isPlaylist {
		p.command = exec.Command(p.name, "-quiet", "-playlist", url)
	} else {
		p.command = exec.Command(p.name, "-quiet", url)
	}

	p.in, err = p.command.StdinPipe()
	if err != nil {
		return err
	}

	p.out, err = p.command.StdoutPipe()
	if err != nil {
		return err
	}

	err = p.command.Start()
	if err != nil {
		return err
	}

	go func() {
		select {
		case <-p.done:
			return
		default:
			p.pipe <- p.out
		}
	}()

	return nil
}

func (p *MPlayer) Mute() {
	p.in.Write([]byte("m"))
}

func (p *MPlayer) Pause() {
	p.in.Write([]byte("p"))
}

func (p *MPlayer) IncVolume() {
	p.in.Write([]byte("*"))
}

func (p *MPlayer) DecVolume() {
	p.in.Write([]byte("/"))
}

func (p *MPlayer) Stop() {
	p.in.Write([]byte("q"))
}

func (p *MPlayer) Close() {
	p.done <- struct{}{}

	p.in.Close()
	p.out.Close()
	close(p.pipe)
}

func getExecutable() (exec string) {
	switch runtime.GOOS {
	case "linux", "darwin":
		exec = "mplayer"
	case "windows":
		exec = "mplayer.exe"
	default:
		exec = "mplayer"
	}

	return
}
