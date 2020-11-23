package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha":34,"bravo":56,"charlie":23,"delta":87}
)

// error but i not find the bug...
func main(){
	fmt.Println("unsorted :")
	for k, v := range barVal {
		fmt.Printf("Key : %v, Value : %v /", k, v)
	}
	// if you wanna sort the map you need make a [] and copy the key into it.
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++ 
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys{
		fmt.Printf("Key: %v, Value: %v /",k,barVal[k])
	}
}
