package main

import "fmt"


var num int = 10
var numx2, numx3 int

func main(){
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2,5,6))

	numx2,numx3 = getX2AndX3(num)
	PrintValues()
	numx2,numx3 = getX2AndX3_2(num)
	PrintValues()
}

func MultiPly3Nums(a int,b int,c int) int {
	return a * b * c
}

func PrintValues(){
	fmt.Printf("num = %d, 2 x num = %d, 3 x num = %d\n",num, numx2, numx3)
}

func getX2andX3(input int) (int, int){
	return 2 * input, 3* input
}

func getX2AndX3_2(input int)(x2 int, x3 int){
	x2 = 2 * input
	x3 = 3 * input
	return
}

