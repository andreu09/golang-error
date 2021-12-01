package main

import (
	"fmt"
	"os"
)

func main() {

	var valEnv string

	if os.Getenv("NAME") != "" {
		valEnv = os.Getenv("NAME")
		fmt.Print(valEnv)
	} else {
		panic("Переменной окружения NAME не существует!")
	}
}
