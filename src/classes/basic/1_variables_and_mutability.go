package basic

import (
	"github.com/maurodesouza/learning-go/src/utils"
)

const IMMUTABLE_CONST = "immutable"              // implicitly typed
const IMMUTABLE_CONST_TYPED string = "immutable" // explicitly typed

var (
	globalStringVar string
	globalIntVar    int
	globalBoolVar   bool
	globalFloatVar  float64
)

// test := "test" // := only works inside a function, so it can't be used at package level

func Class1() {
	// IMMUTABLE_CONST = "mutable" // This will cause a compile error

	utils.PrintHeader("DECLARATION")

	/**
	 *  There are only 3 ways to declare variables in Go:
	 *  1. const name = value
	 *  2. var name = value
	 *  3. name := value
	 *
	 *  Every variable declaration in Go has a initial value
	 *  If not explicitly set, it will be the zero value of the type
	 *
	 *  Zero values:
	 *  - string: ""
	 *  - int: 0
	 *  - bool: false
	 *  - float64: 0.0
	 *  - rune: 0
	 *  - byte: 0
	 */

	{
		utils.PrintSubHeader("Case 1")

		var localStringVar string // Zero value is empty string
		var localIntVar int       // Zero value is 0
		var localBoolVar bool     // Zero value is false
		var localFloatVar float64 // Zero value is 0.0

		var localRune rune = 'a'
		var localByte byte = 'b'

		// Take care about '' and "" - they are different
		// 'a' is a rune (32-bit Unicode code point)
		// "a" is a string (sequence of bytes)

		println(localStringVar)
		println(localIntVar)
		println(localBoolVar)
		println(localFloatVar)
		println(localRune)
		println(localByte)

		utils.PrintSeparator()
	}

	// It's possible to create multiple variables in a single line
	// This is useful for grouping related variables together

	{
		utils.PrintSubHeader("Case 2")

		var (
			a string
			b int
			c float64
			d bool
		)

		println(a)
		println(b)
		println(c)
		println(d)

		utils.PrintSeparator()
	}

	// Using `var name = value`, Go infers the type from the initial value
	// (no explicit type needed). This is a common and idiomatic pattern

	{
		utils.PrintSubHeader("Case 3")

		var (
			a = "string"
			b = 2
			c = 3.0
			d = true
		)

		println(a)
		println(b)
		println(c)
		println(d)

		utils.PrintSeparator()
	}

	// The short declaration `name := value` also infers the type,
	// but unlike `var`, it can only be used inside a function

	{
		utils.PrintSubHeader("Case 4")

		a := "string"
		b := 2
		c := 3.0
		d := true

		println(a)
		println(b)
		println(c)
		println(d)

		utils.PrintSeparator()
	}

	// MUTABILITY

	utils.PrintHeader("MUTABILITY")

	{
		utils.PrintSubHeader("Case 1")

		a := "initial"
		var b = 1

		println(a)
		println(b)

		a = "modified"
		b = 10

		println(a)
		println(b)

		utils.PrintSeparator()
	}

	// := cannot be used to redeclare a variable

	{
		utils.PrintSubHeader("Case 2")

		a := "test"
		// a := "test2"
		println(a)

		utils.PrintSeparator()
	}

	// := can reassign existing variables as long as at least one new variable
	// is introduced on the left side (here `b` is new, so `a` is reassigned)

	{
		utils.PrintSubHeader("Case 3")

		a := "test"
		a, b := "a modified", 2

		println(a)
		println(b)

		utils.PrintSeparator()
	}

	// #region Numbers
	//
	// Unsigned integers
	//
	// | Type    | Bits            | Zero value | Min | Max                  |
	// |---------|-----------------|------------|-----|----------------------|
	// | uint    | 32 or 64 (arch) | 0          | 0   | 2^32-1 or 2^64-1     |
	// | uint8   | 8               | 0          | 0   | 255                  |
	// | uint16  | 16              | 0          | 0   | 65535                |
	// | uint32  | 32              | 0          | 0   | 4294967295           |
	// | uint64  | 64              | 0          | 0   | 18446744073709551615 |
	// | uintptr | 32 or 64 (arch) | 0          | 0   | (enough for pointer) |
	//
	// Aliases: byte = uint8
	//
	// =======================================================================
	//
	// Signed integers
	//
	// | Type  | Bits            | Zero value | Min                  | Max                  |
	// |-------|-----------------|------------|----------------------|----------------------|
	// | int   | 32 or 64 (arch) | 0          | -2^31 or -2^63       | 2^31-1 or 2^63-1     |
	// | int8  | 8               | 0          | -128                 | 127                  |
	// | int16 | 16              | 0          | -32768               | 32767                |
	// | int32 | 32              | 0          | -2147483648          | 2147483647           |
	// | int64 | 64              | 0          | -9223372036854775808 | 9223372036854775807  |
	//
	// Aliases: rune = int32
	//
	// =======================================================================
	//
	// Floating point
	//
	// | Type       | Bits | Zero value | Min (approx)         | Max (approx)         |
	// |------------|------|------------|----------------------|----------------------|
	// | float32    | 32   | 0.0        | ~1.18e-38            | ~3.4e+38             |
	// | float64    | 64   | 0.0        | ~2.23e-308           | ~1.8e+308            |
	//
	// =======================================================================
	//
	// Complex numbers
	//
	// | Type       | Bits | Zero value  | Parts                            |
	// |------------|------|-------------|----------------------------------|
	// | complex64  | 64   | (0+0i)      | float32 real + float32 imaginary |
	// | complex128 | 128  | (0+0i)      | float64 real + float64 imaginary |
	//
	// #endregion
}
