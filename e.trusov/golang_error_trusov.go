package main

import (
	"fmt"
	"os"
) 

func main() {

	name, err := os.LookupEnv("NAME")

	if err == true {
		fmt.Println("Переменная найдена", name)
		return
	}
	panic("Переменная не найдена")

}
