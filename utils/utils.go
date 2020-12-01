package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

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
