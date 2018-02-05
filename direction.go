package main

type element interface {
	direction()
	goLeft()
	goRight()
	goUp()
	goDown()
}
