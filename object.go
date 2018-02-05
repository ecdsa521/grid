package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type object struct {
	left  float64
	top   float64
	size  float64
	color color.RGBA
	image *imdraw.IMDraw
}

func (o *object) draw() {

	//var size float64 = 50
	//o.sprite.Draw(win, pixel.IM.Moved(pixel.Vec{X: o.left * size, Y: o.top * size}))
	o.image.Draw(batch)
}

//phone keyboard directions
func (o *object) move(direction int) {
	switch direction {
	case 1:
		o.left--
		o.top--
	case 2:
		o.top--
	case 3:
		o.left++
		o.top--
	case 4:
		o.left--
	case 6:
		o.left++
	case 7:
		o.left--
		o.top++
	case 8:
		o.top++
	case 9:
		o.left++
		o.top++
	}
	o.pos()
}

func (o *object) pos() {
	x1, x2, y1, y2 := 50+(o.left)*o.size, 50+(o.left+1)*o.size, 750-(o.top)*o.size, 750-(o.top+1)*o.size
	o.image.Clear()
	o.image.Color = o.color
	o.image.Push(pixel.V(x1+1, y1), pixel.V(x1, y2))
	o.image.Push(pixel.V(x2, y2), pixel.V(x2, y1))
	o.image.Push(pixel.V(x2, y1), pixel.V(x1, y1))
	o.image.Line(1)
}
func newObject(left float64, top float64, size float64, color color.RGBA) *object {
	t := new(object)
	t.size = size
	t.left = left
	t.top = top
	t.color = color
	t.image = imdraw.New(nil)
	t.pos()
	return t

}
func newBorder(x1 float64, x2 float64, y1 float64, y2 float64, color color.RGBA) *object {
	t := new(object)
	t.left = x1
	t.top = y1
	t.image = imdraw.New(nil)
	t.image.Color = color
	t.image.Push(pixel.V(x1, y1), pixel.V(x1, y2))
	t.image.Push(pixel.V(x2, y2), pixel.V(x2, y1))
	t.image.Push(pixel.V(x2, y1), pixel.V(x1, y1))
	t.image.Line(1)
	return t
}
func (o *object) playerMove(key pixelgl.Button) {
	switch key {

	case pixelgl.KeyLeft:
		o.left--
		o.pos()
	case pixelgl.KeyRight:
		o.left++
		o.pos()
	case pixelgl.KeyUp:
		o.top--
		o.pos()
	case pixelgl.KeyDown:
		o.top++
		o.pos()
	}

}
