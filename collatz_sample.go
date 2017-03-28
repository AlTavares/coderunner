package main

import (
	"log"
	"os"
	"strconv"
	"fmt"
)

func main() {
	limit, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		log.Fatal(err)
	}
	num, count := maxCollatz(limit)
	fmt.Println(num, count)
}

func maxCollatz(limit int) (maxNum int, maxCount int) {
	var count int
	for i := 0; i < limit; i++ {
		count = collatz(i)
		if count > maxCount {
			maxCount = count
			maxNum = i
		}
	}
	return
}

func collatz(number int) int {
	count := 1
	for number > 1 {
		if number%2 != 0 {
			number = 3*number + 1
			count++
		}
		number = number / 2
		count++
	}
	return count
}
