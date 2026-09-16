package basic

import (
	"fmt"
	"time"

	"github.com/maurodesouza/learning-go/src/utils"
)

func LessonLoops() {

	utils.PrintHeader("LOOPS")
	utils.PrintSeparator()
	utils.PrintSubHeader("While loop")

	i := 0
	for i < 5 {
		i++

		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("For loop")

	for j := 0; j < 5; j += 1 {
		fmt.Println(j + 1)
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Range loop")
	utils.PrintSubHeader("Slice")

	slice := []string{"a", "b", "c", "d", "e"}
	for i, v := range slice {
		fmt.Println(i, v)
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Map")

	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("String")

	for i, r := range "hello" {
		fmt.Println(i, string(r))
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Number")

	for value := range 5 {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 500)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Infinite")

	mark := 0

	for {
		mark++
		fmt.Println("Infinite loop")
		time.Sleep(time.Millisecond * 500)
		if mark == 5 {
			break
		}
	}
}
