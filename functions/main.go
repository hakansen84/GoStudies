package main

import "fmt"

func returnFunc(x string) func(){
	return func(){ fmt.Println(x)}
}
