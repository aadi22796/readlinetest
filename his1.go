package main

import (
    "fmt"
    "os"
    	"bufio"

	"sync"
	"debug/elf"
)
var mutex sync.Mutex{}

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

func fopen(filename string, mode int, perm os.FileMode){
//	if (read mode)
//{
// 	mutex.RLock()
// }	else {
// 	mutex.Lock()
// }

	f, err := os.OpenFile(filename, mode, perm)
	if err!=nil{m
		panic(err)
	}

	//lockerr:=syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	//fmt.Println("File is now locked,")
	//if lockerr!=nil{
	//	panic(lockerr)
	//}
	defer f.Close()

	write(f)

	//lockerr2:=syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	//fmt.Println("File is now unlocked,")
	//if lockerr2!=nil{
	//	panic(lockerr2)
	//}
}

func main() {
	fopen("data", os.O_APPEND|os.O_WRONLY, 0644)
}

//f, err := os.OpenFile("data", O_APPEND|O_WRONLY, 0644)