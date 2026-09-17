package main

import (
	"fmt"

	"github.com/maurodesouza/learning-go/src/classes/basic"
)

const CURRENT_CLASS = "funcs"

func init() {
	fmt.Println("initializing package")
}

func main() {
	switch CURRENT_CLASS {
	case "loops":
		basic.LessonLoops()
	case "variables_and_mutability":
		basic.LessonVariablesAndMutability()
	case "funcs":
		basic.LessonFuncs()
	}
}
