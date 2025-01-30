package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите данные: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Вы ввели следующие данные: %s\n", input)
}
