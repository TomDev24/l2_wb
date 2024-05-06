package main

import "fmt"

// Subsystems
type AudioPlayer struct {
}

func (a *AudioPlayer) PlayAudio() {
	fmt.Println("Playing audio...")
}

type VideoPlayer struct {
}

func (v *VideoPlayer) PlayVideo() {
	fmt.Println("Playing video...")
}

type ScreenManager struct {
}

func (s *ScreenManager) ShowScreen() {
	fmt.Println("Showing screen...")
}

// Facade
type MultimediaFacade struct {
	audioPlayer   *AudioPlayer
	videoPlayer   *VideoPlayer
	screenManager *ScreenManager
}

func NewMultimediaFacade() *MultimediaFacade {
	return &MultimediaFacade{
		audioPlayer:   &AudioPlayer{},
		videoPlayer:   &VideoPlayer{},
		screenManager: &ScreenManager{},
	}
}

func (m *MultimediaFacade) PlayMovie() {
	m.audioPlayer.PlayAudio()
	m.videoPlayer.PlayVideo()
	m.screenManager.ShowScreen()
}

func main() {
	multimediaSystem := NewMultimediaFacade()

	fmt.Println("Playing a movie...")
	multimediaSystem.PlayMovie()
}
