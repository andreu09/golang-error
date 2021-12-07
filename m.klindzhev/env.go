package main

import (
	"fmt"
	"os"
)

func main() {
	envar, exist := os.LookupEnv("NAME")
	if !exist {
		panic("Переменная окружения NAME отсутствует")
	} else {
		fmt.Println("NAME", envar)
	}
}
