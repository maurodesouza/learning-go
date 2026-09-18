package main

import (
	"fmt"

	"github.com/maurodesouza/learning-go/src/classes/basic"
	"github.com/maurodesouza/learning-go/src/classes/goroutines"
)

const CURRENT_CLASS = "wait_group"

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
	case "wait_group":
		goroutines.LessonWaitGroup()
	}
}
