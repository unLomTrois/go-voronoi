package main

import (
	"fmt"
	"math/rand"
	"template/internal/app/geometry"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "evo",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	// generate points
	points := make([]*geometry.Point, 0)

	points = append(points, geometry.NewPoint(win.Bounds(), pixel.V(100, 100)))
	points = append(points, geometry.NewPoint(win.Bounds(), pixel.V(150, 400)))
	points = append(points, geometry.NewPoint(win.Bounds(), pixel.V(400, 300)))

	var (
		frames = 0
		second = time.Tick(time.Second)
		last   = time.Now()
	)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		for _, p := range points {
			p.Draw(imd)
		}

		imd.Draw(win)

		// fps
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("FPS: %d, Delta: %f", frames, dt))
			frames = 0
		default:
		}

		win.Update()
	}
}
