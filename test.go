package main

import (
	"fmt"

	"github.com/johnsudaar/fortunes/picker"
	"github.com/johnsudaar/fortunes/reader"
)

func main() {
	picker, err := picker.LoadPicker("fortunes.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %v fortunes\n", len(picker.Fortunes))

	reader.Read(picker.Pick())
}
