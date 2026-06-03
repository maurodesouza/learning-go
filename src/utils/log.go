package utils

import "fmt"

func PrintHeader(header string) {
	fmt.Println("### " + header + " ###")
}

func PrintSubHeader(subHeader string) {
	fmt.Println("# " + subHeader)
}

func PrintSeparator() {
	fmt.Println("=============================================================")
}
