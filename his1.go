package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
	"syscall"
)

func fwrite(fil *os.File) {
	rd := bufio.NewReader(os.Stdin)
	inp, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if inp == "exit\n" {
	} else {
		circbuf(fil, inp)
		fwrite(fil)
	}
	return
}

func circbuf(fil *os.File, inp string) {
	//creating lastline file if not exists
	if _, err := os.Stat("lastline"); os.IsNotExist(err) {
		f, _ := os.OpenFile("lastline", os.O_WRONLY|os.O_CREATE, 0600)
		f.WriteString("0\n")
	}
	const size = 20
	//reading lastline file and storing in int
	f, err := os.OpenFile("lastline", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	bf := bufio.NewReader(f)
	ll, _ := bf.ReadString('\n')
	lastline, _ := strconv.Atoi(strings.Trim(ll, "\n"))

	//conditions for circular insertion
	switch con := lastline / size; con {
	case 0:
		{
			if lastline%size == size-1 {
				fil.WriteString(strings.Trim(inp, "\n"))
				lastline++
				ioutil.WriteFile("lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)
			} else {
				fil.WriteString(inp)
				lastline++
				ioutil.WriteFile("lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)
			}
		}
	default:
		{
			var n = lastline % size
			d, _ := ioutil.ReadFile("data")
			lines := strings.Split(string(d), "\n")
			lines[n] = strings.Trim(inp, "\n")
			output := strings.Join(lines, "\n")
			ioutil.WriteFile("data", []byte(output), 0600)
			lastline++
			ioutil.WriteFile("lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)
		}
	}
}

func fopen(filename string, mode int, perm os.FileMode) {
	f, err := os.OpenFile(filename, mode, perm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lockerr := syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	if lockerr != nil {
		panic(lockerr)
	}
	fwrite(f)
	unlockerr := syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	if unlockerr != nil {
		panic(lockerr)
	}
}

func main() {
	fopen("data", os.O_APPEND|os.O_WRONLY, 0644)
}