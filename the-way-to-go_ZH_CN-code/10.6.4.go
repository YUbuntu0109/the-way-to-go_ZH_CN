package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct{
	Name string
	Log
}

func main(){
	c := &Customer{"GoogTech", Log{"Yes I am"}}
	c.Add("add a log information")
	fmt.Println(c)
}

func (l *Log) Add(s string) { l.msg += "\n" + s}

func (l *Log) String() string { return l.msg}

func (c *Customer) String() string { return c.Name + "\nLog:" + fmt.Sprintln(c.Log)}


