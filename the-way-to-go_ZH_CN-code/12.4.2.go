package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	inputFile := "files/input.dat"
	outputFile := "files/input_copy.dat"
	buf,err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error : %s\n",err)
	}
	fmt.Printf("%s\n",string(buf))
	err = ioutil.WriteFile(outputFile,buf,1024)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("copy the data from %s file to %s file successfully", inputFile, outputFile)
}

