package main

import "os"

func main(){
	// input data to terminal
	os.Stdout.WriteString("hello world\n")
	f,_ := os.OpenFile("files/write.dat",os.O_CREATE|os.O_WRONLY,0666)
	defer f.Close()
	// input data to a specified file
	f.WriteString("hello world in a file")
}

