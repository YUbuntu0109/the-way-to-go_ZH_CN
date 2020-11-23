package main

import (
	"./person"
	"fmt"
)

func main(){
	p := new(person.Person)
	p.SetFirstName("GoogTech")
	fmt.Println(p.FirstName())
}

