package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errros reading, exiting program..")
		return
	}

	fmt.Printf("Your name is : %s",input)

	// what happended it's not be execute when i input the name is 'Chris' ?
	switch input{
	case "Philip\r\n": fallthrough
	case "Ivo\r\n": fallthrough
	case "Chris\r\n": fmt.Printf("Weclome %s\n", input)
	default: fmt.Printf("You are not weclome here ! GoodBye !\n")
	}
}

