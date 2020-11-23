package main

import "fmt"

type Camera struct{}

func (c *Camera) TackAPicture()string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string{
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main(){
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple beaviors...")
	fmt.Println("It'exhibits behavior of a Camera:",cp.TackAPicture())
	fmt.Println("It works like a Phon too:",cp.Call())

}

