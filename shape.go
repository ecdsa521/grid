package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel/pixelgl"
)

type shape struct {
	rows     int
	cols     int
	left     float64
	size     float64
	top      float64
	rotation int

	grid [][]*object
}

func (s *shape) makeGrid(rows int, cols int) {
	s.grid = make([][]*object, 0)

}

func newShape(rows int, cols int, left float64, top float64, size float64) *shape {
	s := new(shape)
	s.rows = rows
	s.cols = cols
	s.left = left
	s.top = top
	s.size = size
	s.makeGrid(rows, cols)
	return s
}
func (s *shape) rotate() {
	s.transpose()
}

func (s *shape) flipV() {
	newGrid := make([][]*object, 0)
	fmt.Printf("%v\n", s.rows)
	for i := s.rows - 1; i >= 0; i-- {
		newGrid = append(newGrid, s.grid[i])
	}
	s.grid = newGrid
	s.pos()
	s.draw()

}
func (s *shape) transpose() {
	newCols := s.rows
	newRows := s.cols
	newGrid := make([][]*object, newRows)
	for i := range newGrid {
		newGrid[i] = make([]*object, newCols)
	}

	for y := range s.grid {
		for x := range s.grid[y] {
			newGrid[x][y] = s.grid[y][x]
			if newGrid[x][y] != nil {
				newLeft := newGrid[x][y].top
				newTop := newGrid[x][y].left
				newGrid[x][y].left = newLeft
				newGrid[x][y].top = newTop
				newGrid[x][y].pos()
			}
		}

	}

	fmt.Printf("%v %v     %v %v\n", len(s.grid), len(newGrid), len(s.grid[0]), len(newGrid[0]))
	s.rows = newRows
	s.cols = newCols
	s.grid = newGrid
	s.pos()
}
func (s *shape) makeRow(row []*color.RGBA) []*object {
	data := make([]*object, s.cols)
	for i, x := range row {
		if x == nil {
			data[i] = nil
		} else {
			data[i] = newObject(s.left+float64(i), s.top+float64(len(s.grid)), s.size, *x)
		}
	}
	return data
}

//recalculate position of all objects
func (s *shape) pos() {
	for y := range s.grid {
		for x := range s.grid[y] {
			if s.grid[y][x] != nil {
				fmt.Printf("%v %v     ", s.grid[y][x].left, s.grid[y][x].top)
				s.grid[y][x].left = s.left + float64(x)
				s.grid[y][x].top = s.top + float64(y)
				s.grid[y][x].pos()
				fmt.Printf("%v %v \n", s.grid[y][x].left, s.grid[y][x].top)
			}
		}
	}
}
func (s *shape) addRow(row []*color.RGBA) {
	data := s.makeRow(row)
	s.grid = append(s.grid, data)
}

func (s *shape) draw() {
	for _, y := range s.grid {
		for _, x := range y {
			if x != nil {
				x.draw()
			}
		}

	}
}
func (s *shape) moveDirection(direction int) {
	switch direction {
	case 1:
		s.left--
		s.top--
	case 2:
		s.top--
	case 3:
		s.left++
		s.top--
	case 4:
		s.left--
	case 6:
		s.left++
	case 7:
		s.left--
		s.top++
	case 8:
		s.top++
	case 9:
		s.left++
		s.top++
	}
	s.pos()
}

func (s *shape) move(direction int) {
	//
	for _, y := range s.grid {
		for _, x := range y {
			if x != nil {

				x.move(direction)
			}
		}
	}
	s.moveDirection(direction)
}
func (s *shape) playerMove(key pixelgl.Button) {
	switch key {

	case pixelgl.KeyLeft:
		s.move(4)
	case pixelgl.KeyRight:
		s.move(6)
	case pixelgl.KeyUp:
		s.move(2)
	case pixelgl.KeyDown:
		s.move(8)
	case pixelgl.KeyW:
		s.flipV()
	case pixelgl.KeyS:
		s.transpose()
	case pixelgl.KeyA:
		s.flipV()
		s.transpose()
		s.flipV()
	}

}
