package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

func readNumbers(fileName string) []int {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)

	}

	return numbers

}

func Part1() int {
	for i, firstloop := range readNumbers("../inputs.txt") {
		for j, secondloop := range readNumbers("../inputs.txt") {

			if i == j {
				continue
			}
			if firstloop+secondloop == 2020 {
				return firstloop * secondloop
			}
		}

	}

	return 0

}

func Part2() int {
	for i, firstloop := range readNumbers("../inputs.txt") {
		for j, secondloop := range readNumbers("../inputs.txt") {
			for k, thrirdloop := range readNumbers("../inputs.txt") {

				if i == j || i == k || k == j {
					continue
				}
				if firstloop+secondloop+thrirdloop == 2020 {
					return firstloop * secondloop * thrirdloop
				}
			}
		}

	}

	return 0

}

func main() {
	fmt.Println(Part1())
	//	fmt.Println(Part2())

	numbers := readNumbers("../inputs.txt")

	l := len(readNumbers("../inputs.txt"))
	k := 3

	gen := combin.NewCombinationGenerator(l, k)
	fmt.Println("What is this???", gen)
	for gen.Next() {
		s := 0
		p := 1
		for _, i := range gen.Combination(nil) {
			s += numbers[i]
			p *= numbers[i]
		}
		if s == 2020 {
			fmt.Println(p)
		}

	}
}
