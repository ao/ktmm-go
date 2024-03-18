package main

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for {
		time.Sleep(60 * time.Second)
		x, y := robotgo.GetMousePos()
		robotgo.MoveMouse(x+1, y+1)
		time.Sleep(6 * time.Millisecond)
		x2, y2 := robotgo.GetMousePos()
		robotgo.MoveMouse(x2-1, y2-1)
	}
}
