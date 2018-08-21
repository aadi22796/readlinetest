package main

import (
	a "github.com/readlinetest/his1"
	"os"
	"bufio"
	"fmt"
)

func main2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	a.Fopen(text)
	main2()
}

func main(){
	main2()
}