package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите целое число: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Вы ввели число: %s\n", input)
}
