package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field bool "An important answer"
	field2 string "The name of the thing"
	field3 int "How much there are"
}

func main(){
	tt := TagType{true, "GoogTech", 1}
	for i := 0; i < 3; i++{
		refTag(tt,1)
	}
}

func refTag(tt TagType, ix int){
	ttType := reflect.Typeof(tt)
	ixFiled := ttType.Field(ix)
	fmt.Printf("%v\n",ixField.Tag)
}

