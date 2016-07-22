//  Выводим индекс и значение каждого аргумента по одному аргументу в строке.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(s)
}
