package facade

import (
	"fmt"
)

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is turned on.")
}

func (t *TV) Off() {
	fmt.Println("TV is turned off.")
}

type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound system is turned on.")
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound system is turned off.")
}

type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD player is turned on.")
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD player is turned off.")
}

// Facade
type HomeTheaterFacade struct {
	tv          *TV
	soundSystem *SoundSystem
	dvdPlayer   *DVDPlayer
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:          &TV{},
		soundSystem: &SoundSystem{},
		dvdPlayer:   &DVDPlayer{},
	}
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("Get ready to watch a movie...")
	h.tv.On()
	h.soundSystem.On()
	h.dvdPlayer.On()
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down the home theater system...")
	h.tv.Off()
	h.soundSystem.Off()
	h.dvdPlayer.Off()
}
