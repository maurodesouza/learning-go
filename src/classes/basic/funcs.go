package basic

import "github.com/maurodesouza/learning-go/src/utils"

func LessonFuncs() {

	utils.PrintHeader("FUNCS")
	utils.PrintSeparator()
	utils.PrintSubHeader("Immediately invoked (IIFE)")

	{
		result := func(a int, b int) int {
			return a + b
		}(1, 2)

		println(result)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Assigned to variable")

	{
		sum := func(a int, b int) int {
			return a + b
		}

		result := sum(1, 2)
		println(result)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Combined param types")

	{
		sum := func(a, b int) int {
			return a + b
		}

		result := sum(1, 2)
		println(result)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Variadic params")

	{
		sum := func(numbers ...int) int {
			total := 0
			for _, num := range numbers {
				total += num
			}
			return total
		}

		result := sum(1, 2, 3, 4, 5)
		println(result)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Multiple returns")

	{
		calc := func(numbers ...int) (int, int) {
			sumTotal := 0
			for _, num := range numbers {
				sumTotal += num
			}

			multiplyTotal := 1
			for _, num := range numbers {
				multiplyTotal *= num
			}

			return sumTotal, multiplyTotal
		}

		sum, multiply := calc(1, 2, 3, 4, 5)
		println(sum, multiply)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Named returns")

	{
		calc := func(numbers ...int) (sumTotal int, multiplyTotal int) {
			sumTotal = 0
			for _, num := range numbers {
				sumTotal += num
			}

			multiplyTotal = 1
			for _, num := range numbers {
				multiplyTotal *= num
			}

			return
		}

		slice := []int{1, 2, 3, 4, 5}

		sum, multiply := calc(slice...)
		println(sum, multiply)
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Closure")

	{
		counter := func() func() int {
			counter := 0

			return func() int {
				counter++
				return counter
			}
		}

		c := counter()
		println(c())
		println(c())
		println(c())
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Function as argument")

	{
		counter := func(increment func() int) func() int {
			counter := 0

			return func() int {
				counter += increment()
				return counter
			}
		}

		c := counter(func() int {
			return 2
		})

		println(c())
		println(c())
		println(c())
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Higher-order function")

	{
		counter := func(increment func() int) func() int {
			counter := 0

			return func() int {
				counter += increment()
				return counter
			}
		}

		c := counter(func() int {
			return 2
		})

		println(c())
		println(c())
		println(c())
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Defer")

	{
		first := func() { println("I'm first") }
		second := func() { println("I'm second") }

		defer first()
		second()
	}

	utils.PrintSeparator()
	utils.PrintSubHeader("Defer with return")

	{
		deferWithReturn := func() string {
			defer println("before return")
			println("Initialization")
			return "message"
		}
		println(deferWithReturn())
	}
}
