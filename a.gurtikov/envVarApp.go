package main

import (
	"fmt"
	"os"
)

func main() {
	name, ok := os.LookupEnv("NAME")
	if !ok {
		panic("переменная NAME не найдена!")
	} else {
		fmt.Println("Переменная NAME ", name)
	}

}
