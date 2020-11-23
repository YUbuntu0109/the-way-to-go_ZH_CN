package main

import "fmt"

type InVector []int

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int{
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func main(){
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is : %d\n",two1.AddThem())
	fmt.Printf("Add them to the param : %d\n", two1.AddToParam(20))

	two2 := TwoInts{3,4}
	fmt.Printf("The sum is : %d\n",two2.AddThem())

	fmt.Println(InVector{1,2,3}.Sum())
}

func (v InVector) Sum()(s int) {
	for _,x := range v{
		s += x
	}
	return
}

