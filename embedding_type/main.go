package main

import "fmt"

type Ball struct {
	Radius int
}

func (b Ball) Bounce() {
	fmt.Println(b, "has radius", b.Radius)
}

type BasketBall struct {
	Ball
	Weight int
}

type FootBall struct {
	Ball
	Weight int
	Radius int
}

type BaseBall struct {
	*Ball
	Weight int
}

type VolleyBall struct {
	Action
	Weight int
}

type Action interface {
	Bounce()
}

func BounceIt(b Action) {
	b.Bounce()
}

func BounceBall(b Ball) {
	b.Bounce()
}

func main() {
	bb := BasketBall{Ball{Radius: 5}, 50}
	bb.Bounce()

	BounceIt(bb)

	//cannot casting embedding to embedded
	// var i interface{}
	// i = bb
	// BounceBall(i.(Ball))

	fb := FootBall{Ball{Radius: 5}, 50, 10}
	fb.Bounce()
	
	baseball := BaseBall{&Ball{5}, 60}
	baseball.Bounce()
	fmt.Println("access radius",baseball.Radius)


	vb := VolleyBall{&Ball{5}, 60}
	vb.Bounce()
	BounceIt(vb)
	// cannot access interface embedded properties
	// fmt.Println("access radius",vb.Radius)
}
