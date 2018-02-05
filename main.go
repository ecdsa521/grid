package main

import (
	"fmt"
	_ "image"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var win *pixelgl.Window
var batch *pixel.Batch

func draw(win *pixelgl.Window, sprite *pixel.Sprite, x float64, y float64) {
	sprite.Draw(win, pixel.IM.Moved(pixel.Vec{X: x, Y: y}))

}

func main() {

	pixelgl.Run(run)
}
func pollKeys() pixelgl.Button {
	if win.Pressed(pixelgl.KeyLeft) {
		return pixelgl.KeyLeft
	}
	if win.Pressed(pixelgl.KeyRight) {
		return pixelgl.KeyRight
	}
	if win.Pressed(pixelgl.KeyUp) {
		return pixelgl.KeyUp
	}
	if win.Pressed(pixelgl.KeyDown) {
		return pixelgl.KeyDown
	}
	if win.Pressed(pixelgl.KeySpace) {
		return pixelgl.KeySpace
	}
	if win.Pressed(pixelgl.KeyEscape) {
		return pixelgl.KeyEscape
	}
	if win.Pressed(pixelgl.KeyW) {
		return pixelgl.KeyW
	}
	if win.Pressed(pixelgl.KeyA) {
		return pixelgl.KeyA
	}
	if win.Pressed(pixelgl.KeyS) {
		return pixelgl.KeyS
	}
	if win.Pressed(pixelgl.KeyD) {
		return pixelgl.KeyD
	}
	return pixelgl.KeyUnknown
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "grid",
		Bounds: pixel.R(0, 0, 1000, 800),
		VSync:  true,
	}
	win, _ = pixelgl.NewWindow(cfg)
	batch = pixel.NewBatch(&pixel.TrianglesData{}, win)
	win.Clear(colornames.Black)

	border := newBorder(50, 950, 50, 750, colornames.White)

	brick := newShape(4, 5, 2, 2, 50, 1)

	red := colornames.Red
	yellow := colornames.Yellow
	brick.addRow2([]*object{
		&object{color: red}, nil, &object{color: yellow, thickness: 1}, nil,
		&object{color: red, thickness: 1, shape: func(o *object) {
			o.image.Ellipse(pixel.V(15, 15), o.thickness+5)
		}, push: func(o *object) {
			x1, x2, y1, y2 := 50+(o.left)*o.size, 50+(o.left+1)*o.size, 750-(o.top)*o.size, 750-(o.top+1)*o.size
			o.image.Push(pixel.V(x1+1, y1), pixel.V(x1, y2))
			o.image.Push(pixel.V(x2, y2), pixel.V(x2, y1))
			o.image.Push(pixel.V(x1+1, y1), pixel.V(x1, y2))
			o.image.Push(pixel.V(x2, y1), pixel.V(x1, y1))

		}},
	})
	//brick.addRow([]*color.RGBA{nil, &colornames.Red, nil, &colornames.Green})
	//brick.addRow([]*color.RGBA{nil, nil, &colornames.Blue, nil, nil})
	//brick.addRow([]*color.RGBA{&colornames.Yellow, nil, nil, nil, &colornames.Yellow})
	//brick.addRow([]*color.RGBA{nil, &colornames.Yellow, &colornames.Yellow, &colornames.Yellow, nil})

	//brick := newShape(3, 3)

	var frames = 0
	var second = time.Tick(time.Second)
	var mil = time.Tick(time.Second / 10)

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()

		frames++
		win.Clear(colornames.Black)
		batch.Clear()
		border.draw()
		brick.draw()

		batch.Draw(win)
		win.Update()
		select {
		case <-mil:
			brick.playerMove(pollKeys())
		case <-second:

			win.SetTitle(fmt.Sprintf("%s | FPS: %d | %v", cfg.Title, frames, dt))
			frames = 0
		default:
		}

	}

}
