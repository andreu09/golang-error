package main
import (
	"fmt"
	"os"
)
func main(){
	name, ok := os.LookupEnv("NAME")
	if ok != true {
		panic("Переменная не найдена")
	}else{
		fmt.Println("Переменная NAME: ", name)
	}
}
