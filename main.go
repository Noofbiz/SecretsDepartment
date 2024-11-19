package main

import (
	"github.com/EngoEngine/engo"
)

type defaultScene struct{}

func (*defaultScene) Type() string { return "Default Scene" }

func (*defaultScene) Preload() {}

func (*defaultScene) Setup(u engo.Updater) {}

func main() {
	engo.Run(engo.RunOptions{
		Title:  "Secrets Department",
		Width:  640,
		Height: 480,
	}, &defaultScene{})
}
