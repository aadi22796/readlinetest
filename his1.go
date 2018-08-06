package main

import (
    "fmt"
    "os"
    	"bufio"
	"sync"
	"syscall"
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

func fopen(filename string, mode int, perm os.FileMode, wg sync.WaitGroup, m sync.Mutex){
	m.Lock()
	f, err := os.OpenFile(filename, mode, perm)
	lockerr:=syscall.Flock(int(f.Fd()), 0600)
	if lockerr!=nil{
		panic(lockerr)
	}
	if err!=nil{
		panic(err)
	}
	defer f.Close()
	write(f)
	m.Unlock()
	wg.Done()
	lockerr2:=syscall.Flock(int(f.Fd()), 0644)
	if lockerr2!=nil{
		panic(lockerr2)
	}
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	w.Add(1)
	go fopen("data", os.O_APPEND|os.O_WRONLY, 0644, w, m)
	w.Wait()
}

//f, err := os.OpenFile("data", O_APPEND|O_WRONLY, 0644)