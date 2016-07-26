// Выводим индекс и значение каждого аргумента по одному аргументу в строке.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
