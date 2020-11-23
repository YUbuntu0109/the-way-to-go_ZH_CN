package main
import (
	"fmt"
	"strings"
)

type Person struct {
	firstname string
	lastname string
}

func upPerson(p *Person){
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToUpper(p.lastname)
}

func main(){

	// first method: as a value type
	var person1 Person
	person1.firstname = "goog"
	person1.lastname = "tech"
	upPerson(&person1)
	fmt.Printf("The name of the person is : %s %s \n",person1.firstname,person1.lastname)

	// second method: as a pointer
	person2 := new(Person)
	person2.firstname = "goog"
	person2.lastname = "tech"
	(*person2).lastname = "tech"
	upPerson(person2)
	fmt.Printf("The name of the person is : %s %s \n",person2.firstname,person2.lastname)

	// third method: as a literal
	person3 := &Person{"goog","tech"}
	upPerson(person3)
	fmt.Printf("The name of the person is %s %s \n",person3.firstname,person3.lastname)
}

