package main

import (
    "fmt"
    "os"
    	"bufio"
	)

func write(fil *os.File){
	fmt.Println("Enter your string")
	rd:=bufio.NewReader(os.Stdin)
	inp, err:=rd.ReadString('\n')
	if err!=nil{
		panic(err)
	}
	if inp == "exit\n"{
		fmt.Println("Exited program")
	} else {
		fil.WriteString(inp)
		write(fil)
	}
	return
}
func main() {
	f, err := os.OpenFile("data", os.O_APPEND|os.O_WRONLY, 0644)
	if err!=nil{
		panic(err)
		}
	defer f.Close()
	write(f)
}