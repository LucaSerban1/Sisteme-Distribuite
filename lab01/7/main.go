package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	slice := make([]string, 0)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	
	m := map[string]int{}
	for _, cuv := range slice {
		m[cuv]++
	}
	fmt.Println(m)
}