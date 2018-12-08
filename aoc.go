package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jassler/aoc18/day01"
	"github.com/jassler/aoc18/day02"
)

var functions = map[string]func(input string, ch chan<- string){
	"1": day01.Start,
	"2": day02.Start,
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Specify day number to solve")
		fmt.Println("Currently available:")
		for day := range functions {
			fmt.Println("   ", day)
		}
		return
	}

	day := os.Args[1]

	inputPath, err := filepath.Abs("input")
	if err != nil {
		panic(err)
	}

	inputFile := path.Join(inputPath, "day"+day+"_input.txt")

	ch := make(chan string, 2)

	functions[day](inputFile, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
