package main

import "github.com/maurodesouza/learning-go/src/classes/basic"

const CURRENT_CLASS = "loops"

func main() {
	switch CURRENT_CLASS {
	case "loops":
		basic.LoopsClass()
	case "variables_and_mutability":
		basic.ClassVariablesAndMutability()
	}
}
